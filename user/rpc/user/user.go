// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package user

import (
	"context"

	"lesjp/user/rpc/rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginReq  = rpc.LoginReq
	LoginResp = rpc.LoginResp

	User interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := rpc.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}
