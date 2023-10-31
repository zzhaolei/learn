package carlogic

import (
	"context"
	"fmt"

	"golang/go-zero/car/client/car"

	"golang/go-zero/gateway/api/car/v1"
	"golang/go-zero/gateway/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *v1.Request) (*v1.Response, error) {
	l.Info("gateway: call ping func in gateway server")

	response, err := l.svcCtx.Config.Service.CarClient.Ping(l.ctx, &car.Request{Ping: in.Ping})
	if err != nil {
		return nil, fmt.Errorf("call car service err %s", err)
	}
	return &v1.Response{
		Pong: response.Pong,
	}, nil
}
