package initialize

import (
	carServer "gozero/car/api/car/v1"
	"gozero/gateway/internal/svc"

	"github.com/zeromicro/go-zero/zrpc"
)

func RegisterService(svcCtx *svc.ServiceContext) {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{"127.0.0.1:8082"},
	})
	// Register car service
	svcCtx.Config.Service.CarClient = carServer.NewCarClient(conn.Conn())
}
