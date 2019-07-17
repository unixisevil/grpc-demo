package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"github.com/unixisevil/grpc-demo/greet/greetpb"
)

func main() {
	tls := flag.Bool("tls", false, "enable tls connect")
	flag.Parse()
	opts := grpc.WithInsecure()
	if *tls {
		certFile := "tls/localhost+2.pem"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatal("loading certFile: %v", err)
		}
		opts = grpc.WithTransportCredentials(creds)
	}
	conn, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal("could not connect: %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	unary(c)
	serverStream(c)
	clientStream(c)
	biStream(c)
	unaryWithDeadLine(c, 1)
	unaryWithDeadLine(c, 5)
}

func unaryWithDeadLine(c greetpb.GreetServiceClient, t time.Duration) {
	req := &greetpb.GreetDeadLineReq{
		Greeting: &greetpb.Greeting{
			FirstName: "jian",
			LastName:  "yu",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), t*time.Second)
	defer cancel()

	res, err := c.GreetWithDeadLine(ctx, req)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			if respErr.Code() == codes.DeadlineExceeded {
				log.Printf("timeout occurred")
			}
		} else {
			log.Printf("we got grpc error: %v\n", err)
		}
		return
	}
	log.Printf("recv GreetWithDeadLine rpc result: %v\n", res.Result)
}

func unary(c greetpb.GreetServiceClient) {
	log.Println("in unary func()")
	req := &greetpb.GreetReq{
		Greeting: &greetpb.Greeting{
			FirstName: "jian",
			LastName:  "yu",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet rpc: %v", err)
	}
	log.Printf("response from greet: %v\n", res.Result)
}

func serverStream(c greetpb.GreetServiceClient) {
	log.Println("in serverStream func()")
	req := &greetpb.GreetManyReq{
		Greeting: &greetpb.Greeting{
			FirstName: "jian",
			LastName:  "yu",
		},
	}
	stream, err := c.GreetMany(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetMany rpc: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while recving stream: %v", err)
		}
		log.Printf("response from GreetMany rpc: %v\n", msg.GetResult())
	}
}

func clientStream(c greetpb.GreetServiceClient) {
	log.Println("in clientStream func()")
	stream, err := c.GreetLong(context.Background())
	if err != nil {
		log.Fatalf("error while calling GreetLong: %v", err)
	}

	reqs := []*greetpb.GreetLongReq{
		&greetpb.GreetLongReq{
			Greeting: &greetpb.Greeting{
				FirstName: "chandler",
				LastName:  "bing",
			},
		},
		&greetpb.GreetLongReq{
			Greeting: &greetpb.Greeting{
				FirstName: "ross",
				LastName:  "geller",
			},
		},
		&greetpb.GreetLongReq{
			Greeting: &greetpb.Greeting{
				FirstName: "joey",
				LastName:  "tribbiani",
			},
		},
	}
	for _, req := range reqs {
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while recving long greet resp: %v", err)
	}
	log.Printf("recv long greet resp:\n%v\n", res.Result)
}

func biStream(c greetpb.GreetServiceClient) {
	log.Println("in biStream func()")
	stream, err := c.GreetEveryOne(context.Background())
	if err != nil {
		log.Fatalf("error while calling GreetEveryOne: %v", err)
	}
	reqs := []*greetpb.GreetEveryOneReq{
		&greetpb.GreetEveryOneReq{
			Greeting: &greetpb.Greeting{
				FirstName: "chandler",
				LastName:  "bing",
			},
		},
		&greetpb.GreetEveryOneReq{
			Greeting: &greetpb.Greeting{
				FirstName: "ross",
				LastName:  "geller",
			},
		},
		&greetpb.GreetEveryOneReq{
			Greeting: &greetpb.Greeting{
				FirstName: "joey",
				LastName:  "tribbiani",
			},
		},
	}
	done := make(chan struct{})
	go func() {
		for _, req := range reqs {
			stream.Send(req)
		}
		stream.CloseSend()
	}()
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("error while recving stream: %v\n", err)
				break
			}
			log.Printf("response from GreetEveryOne rpc: %v\n", msg.GetResult())
		}
		close(done)
	}()
	<-done
}
