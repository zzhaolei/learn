package main

import (
	"context"
	"log"
	"time"

	hello "golang/grpc/hello_w"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("dial err: %s", err)
	}

	c := hello.NewNiHaoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	name := "你好"
	// r, err := c.SayHello(ctx, &hello.HelloRequest{Name: &name})
	r, err := c.SayHelloAgin(ctx, &hello.HelloRequest{Name: &name})
	if err != nil {
		log.Fatalf("hello request err: %s", err)
	}

	log.Printf("receive: %s", r)
}
