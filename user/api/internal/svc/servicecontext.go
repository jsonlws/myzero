package svc

import (
	"lesjp/user/api/internal/config"
	"lesjp/user/api/internal/middleware"
	"lesjp/user/rpc/user"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
	Example rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Example: middleware.NewExampleMiddleware().Handle,
	}
}
