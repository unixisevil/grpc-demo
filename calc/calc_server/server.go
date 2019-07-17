package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/unixisevil/grpc-demo/calc/calcpb"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calcpb.SumReq) (*calcpb.SumRes, error) {
	return &calcpb.SumRes{
		Res: req.FirstNum + req.SecondNum,
	}, nil
}

func (*server) PrimeFactor(req *calcpb.PrimeFactorReq, stream calcpb.CalcService_PrimeFactorServer) error {
	log.Printf("recv PrimeFactor rpc req:%v\n", req)
	num := req.GetNumber()
	divisor := int64(2)
	for num > 1 {
		if num%divisor == 0 {
			stream.Send(&calcpb.PrimeFactorRes{
				PrimeFactor: divisor,
			})
			num /= divisor
		} else {
			divisor++
		}
	}
	return nil
}

func (*server) Avg(stream calcpb.CalcService_AvgServer) error {
	log.Println("recv compute avg rpc")
	sum := int64(0)
	count := int64(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(
				&calcpb.AvgRes{
					Avg: float64(sum) / float64(count),
				},
			)
		}
		if err != nil {
			log.Fatalf("failed to recv client stream: %v", err)
		}
		sum += req.GetNumber()
		count++
	}
}

func (*server) Max(stream calcpb.CalcService_MaxServer) error {
	log.Println("recv compute max rpc")
	max := int64(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("recv end of stream\n")
			return nil
		}
		if err != nil {
			log.Printf("error while max recving stream: %v\n", err)
			return err
		}
		num := req.GetNumber()
		if num > max {
			max = num
			if err = stream.Send(&calcpb.MaxRes{
				Max: max,
			}); err != nil {
				log.Printf("error while max sending stream: %v\n", err)
				return err
			}
		}
	}
}

func (*server) Sqrt(ctx context.Context, req *calcpb.SqrtReq) (*calcpb.SqrtRes, error) {
	num := req.GetNumber()
	if num < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("received negative number:%v", num),
		)
	}
	return &calcpb.SqrtRes{
		Root: math.Sqrt(num),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	calcpb.RegisterCalcServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serv: %v\n", err)
	}
}
