package main

import (
	"flag"
	"fmt"
	"net/http"

	carV1 "gozero/gateway/api/car/v1"
	userV1 "gozero/gateway/api/user/v1"
	"gozero/gateway/docs"
	"gozero/gateway/initialize"

	"github.com/zeromicro/go-zero/rest"

	"gozero/gateway/internal/config"
	carServer "gozero/gateway/internal/server/car"
	userServer "gozero/gateway/internal/server/user"
	"gozero/gateway/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	zGateway "github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/gateway.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	// register service client into `ctx` config
	initialize.RegisterService(ctx)

	// grpc
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		carV1.RegisterCarServer(grpcServer, carServer.NewCarServer(ctx))
		userV1.RegisterUserServer(grpcServer, userServer.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	// http
	gw := zGateway.MustNewServer(c.Gateway)
	gw.AddRoutes([]rest.Route{
		{
			Method: http.MethodGet,
			Path:   "/docs",
			Handler: func(w http.ResponseWriter, req *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write(docs.SwaggerUI)
			},
		},
		{
			Method: http.MethodGet,
			Path:   "/openapi.yaml",
			Handler: func(w http.ResponseWriter, req *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write(docs.OpenAPI)
			},
		},
	})

	// group
	group := service.NewServiceGroup()
	group.Add(s)
	group.Add(gw)
	defer group.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	fmt.Printf("Starting http server at %s:%d...\n", c.Gateway.Host, c.Gateway.Port)
	group.Start()
}
