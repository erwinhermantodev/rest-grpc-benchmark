package grpc

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/erwinhermanto31/service2/transport/grpc/proto/service2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run handler running grpc
func Run() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("grpc_port"))
	if err != nil {
		log.Fatalf("failed to listern :%v", err)
	} else {
		fmt.Printf(" Connect grpc to port :%v \n", os.Getenv("grpc_port"))
	}

	s := grpc.NewServer()
	pb.RegisterService2Server(s, NewGrpcServer())

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
