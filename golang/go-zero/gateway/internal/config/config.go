package config

import (
	"gozero/car/client/car"

	"github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Gateway gateway.GatewayConf

	Service Service `json:",optional"`
}

type Service struct {
	CarClient car.Car
}
