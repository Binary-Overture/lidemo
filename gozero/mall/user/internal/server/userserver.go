// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"
	"firstproject/gozero/mall/user/internal/logic"
	"firstproject/gozero/mall/user/internal/svc"
	user2 "firstproject/gozero/mall/user/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user2.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *user2.IdRequest) (*user2.UserResponse, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}
