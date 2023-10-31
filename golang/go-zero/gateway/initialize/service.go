package initialize

import (
	"github.com/zeromicro/go-zero/zrpc"
	carServer "golang/go-zero/car/api/car/v1"
	"golang/go-zero/gateway/internal/svc"
)

func RegisterService(svcCtx *svc.ServiceContext) {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{"127.0.0.1:8082"},
	})
	// Register car service
	svcCtx.Config.Service.CarClient = carServer.NewCarClient(conn.Conn())
}
