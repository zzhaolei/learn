// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"golang/go-zero/gateway/api/user/v1"
	"golang/go-zero/gateway/internal/logic/user"
	"golang/go-zero/gateway/internal/svc"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	v1.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Ping(ctx context.Context, in *v1.Request) (*v1.Response, error) {
	l := userlogic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}