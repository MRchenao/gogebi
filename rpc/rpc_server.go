package rpc

import (
	"context"
	"gebi/app/Http/Serializer"
	hello "gebi/proto"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
}

const (
	port = ":50051"
)

func RpcService() {
	g := grpc.NewServer()
	s := Server{}
	hello.RegisterGreeterServer(g, &s)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	if err := g.Serve(lis); err != nil {
		Serializer.Err(30333, "fail to server:%#v", err)
	}
}

func (s *Server) SayHello(ctx context.Context, request *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{Message: "Hello " + request.Name}, nil
}

func (s *Server) SayHelloAgain(ctx context.Context, request *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{Message: "Hello Again " + request.Name}, nil
}
