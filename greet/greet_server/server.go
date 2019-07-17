package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"github.com/unixisevil/grpc-demo/greet/greetpb"
)

type server struct{}

func (*server) GreetWithDeadLine(ctx context.Context, req *greetpb.GreetDeadLineReq) (*greetpb.GreetDeadLineRes, error) {
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			log.Println("client cancel request")
			return nil, status.Error(codes.Canceled, "the client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	return &greetpb.GreetDeadLineRes{
		Result: fmt.Sprintf("Hello, %s %s", firstName, lastName),
	}, nil
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetReq) (*greetpb.GreetRes, error) {
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	return &greetpb.GreetRes{
		Result: fmt.Sprintf("Hello, %s %s", firstName, lastName),
	}, nil
}

func (*server) GreetMany(req *greetpb.GreetManyReq, stream greetpb.GreetService_GreetManyServer) error {
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	for i := 0; i < 10; i++ {
		stream.Send(
			&greetpb.GreetManyRes{
				Result: fmt.Sprintf("Hello, %s %s", firstName, lastName),
			},
		)
		time.Sleep(100 * time.Microsecond)
	}
	return nil
}

func (*server) GreetLong(stream greetpb.GreetService_GreetLongServer) error {
	fmt.Println("recving client stream")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.GreetLongRes{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("error while recving client stream: %v\n", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		lastName := req.GetGreeting().GetLastName()
		result += fmt.Sprintf("Hello, %s %s\n", firstName, lastName)
	}
}

func (*server) GreetEveryOne(stream greetpb.GreetService_GreetEveryOneServer) error {
	fmt.Println("bidirectional stream")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("recv stream error: %v\n", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		lastName := req.GetGreeting().GetLastName()
		err = stream.Send(&greetpb.GreetEveryOneRes{
			Result: fmt.Sprintf("hi %s %s !", firstName, lastName),
		})
		if err != nil {
			log.Printf("send stream error: %v\n", err)
			return err
		}
	}
}

func main() {
	tls := flag.Bool("tls", false, "enable tls connect")
	flag.Parse()
	fmt.Println("server started")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	if *tls {
		certFile := "tls/localhost+2.pem"
		keyFile := "tls/localhost+2-key.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("failed to loading certFile: %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	s := grpc.NewServer(opts...)
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serv: %v\n", err)
	}
}
