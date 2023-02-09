package grpc

import (
	"context"

	"github.com/erwinhermanto31/service1/builder"
	pb "github.com/erwinhermanto31/service1/transport/grpc/proto/service2"
)

type GrpcServer struct {
	builder *builder.Grpc
	pb.UnimplementedService2Server
}

func (gs *GrpcServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	return &pb.HelloResponse{Message: "hello :" + req.Name}, nil
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{
		builder: builder.NewGrpc(),
	}
}
