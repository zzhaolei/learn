// Code generated by goctl. DO NOT EDIT.
// Source: car.proto

package car

import (
	"context"

	"gozero/car/api/car/v1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = v1.Request
	Response = v1.Response

	Car interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultCar struct {
		cli zrpc.Client
	}
)

func NewCar(cli zrpc.Client) Car {
	return &defaultCar{
		cli: cli,
	}
}

func (m *defaultCar) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := v1.NewCarClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
