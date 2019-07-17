package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/unixisevil/grpc-demo/calc/calcpb"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect: %v", err)
	}
	defer conn.Close()
	c := calcpb.NewCalcServiceClient(conn)
	unary(c)
	serverStream(c)
	clientStream(c)
	biStream(c)
	unaryWithError(c)
}

func unary(c calcpb.CalcServiceClient) {
	req := &calcpb.SumReq{
		FirstNum:  10,
		SecondNum: 20,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet rpc: %v", err)
	}
	log.Printf("response from greet: %v\n", res.Res)
}

func serverStream(c calcpb.CalcServiceClient) {
	req := &calcpb.PrimeFactorReq{
		Number: 12390392840,
	}
	stream, err := c.PrimeFactor(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeFactor rpc: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while recv server stream: %v", err)
		}
		log.Println(res.PrimeFactor)
	}

}

func clientStream(c calcpb.CalcServiceClient) {
	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("error while calling Avg rpc: %v", err)
	}
	for _, num := range []int64{11, 39, 40, 90, 100, 76} {
		stream.Send(&calcpb.AvgReq{Number: num})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while recv avg result: %v", err)
	}
	log.Printf("avg is %f\n", res.Avg)
}

func biStream(c calcpb.CalcServiceClient) {
	log.Println("biStream client started")
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("error while calling max rpc: %v", err)
	}
	done := make(chan struct{})
	go func() {
		for _, num := range []int64{13, 87, 1, 4, 100, 92, 33} {
			if err := stream.Send(&calcpb.MaxReq{
				Number: num,
			}); err != nil {
				log.Printf("error while max sending stream: %v\n", err)
			}
		}
		stream.CloseSend()
	}()
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("error while recving stream: %v\n", err)
				break
			}
			log.Printf("response from Max rpc: %v\n", res.Max)
		}
		close(done)
	}()
	<-done
}

func unaryWithError(c calcpb.CalcServiceClient) {
	unaryWithErrorCheck(c, 2)
	unaryWithErrorCheck(c, -2)
}

func unaryWithErrorCheck(c calcpb.CalcServiceClient, num float64) {
	req := &calcpb.SqrtReq{Number: num}
	res, err := c.Sqrt(context.Background(), req)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			log.Printf("code: %v, msg: %v\n", respErr.Code(), respErr.Message())
			if respErr.Code() == codes.InvalidArgument {
				log.Printf("we sent a negative number to server")
			}
		} else {
			log.Printf("we got grpc error: %v\n", err)
		}
		return
	}
	log.Printf("sqrt result: %v\n", res.Root)
}
