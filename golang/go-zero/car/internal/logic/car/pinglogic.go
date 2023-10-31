package carlogic

import (
	"context"
	"fmt"

	"golang/go-zero/car/api/car/v1"
	"golang/go-zero/car/internal/svc"

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
	l.Info("car: call ping func in car rpc server.")
	return &v1.Response{
		Pong: fmt.Sprintf("This car server response. Request: %s.", in.Ping),
	}, nil
}
