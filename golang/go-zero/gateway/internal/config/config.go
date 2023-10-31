package config

import (
	"github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/zrpc"
	"golang/go-zero/car/client/car"
)

type Config struct {
	zrpc.RpcServerConf
	Gateway gateway.GatewayConf

	Service Service `json:",optional"`
}

type Service struct {
	CarClient car.Car
}
