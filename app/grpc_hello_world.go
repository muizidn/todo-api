package app

import (
	"context"

	pb "github.com/muizidn/todo-api/pb"
)

type serverGreeter struct {
	pb.UnimplementedGreeterServer
}

func (s *serverGreeter) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
