package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/erwinhermanodev/rest-grpc-benchmark/grpc/proto" // replace with actual path

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	time.Sleep(10 * time.Millisecond) // simulate some processing time
	return &pb.HelloResponse{Message: "Hello, gRPC!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	log.Println("Starting gRPC server on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
