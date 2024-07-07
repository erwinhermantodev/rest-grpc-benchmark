package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "github.com/erwinhermanodev/rest-grpc-benchmark/client/proto" // replace with actual path

	"google.golang.org/grpc"
)

func benchmarkREST() {
	start := time.Now()
	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		log.Fatalf("failed to call REST API: %v", err)
	}
	defer resp.Body.Close()
	duration := time.Since(start)
	fmt.Printf("REST API call took %v\n", duration)
}

func benchmarkGRPC() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	start := time.Now()
	_, err = c.SayHello(ctx, &pb.HelloRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	duration := time.Since(start)
	fmt.Printf("gRPC call took %v\n", duration)
}

func main() {
	for i := 0; i < 10; i++ {
		benchmarkREST()
		benchmarkGRPC()
	}
}
