package main

import (
	"context"
	"log"
	"net"

	hello "golang/grpc/hello_w"

	"google.golang.org/grpc"
)

type server struct {
	hello.UnimplementedNiHaoServer
}

func (s *server) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{
		Message: req.Name,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("listen err: %s", err)
	}

	s := grpc.NewServer()
	hello.RegisterNiHaoServer(s, &server{})
	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
