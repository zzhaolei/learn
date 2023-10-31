package userlogic

import (
	"context"

	"golang/go-zero/gateway/api/user/v1"
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
	_ = in
	return &v1.Response{Pong: "pong"}, nil
}
