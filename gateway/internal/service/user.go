package service

import (
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/zrpc"
	"miya/gateway/internal/config"
	user2 "miya/user/rpc/types/user"
	"miya/user/rpc/user"
)

type ServiceContext struct {
	UserRpc user.User
	Context *gin.Context
}

func NewServiceContext(context *gin.Context) *ServiceContext {
	rpcConfig := config.GetConfig().UserRpc
	client := zrpc.MustNewClient(rpcConfig)
	rp := user.NewUser(client)
	return &ServiceContext{
		UserRpc: rp,
		Context: context,
	}
}

func (l *ServiceContext) GetUser() string {
	//@todo test
	user, err := l.UserRpc.GetUser(l.Context, &user2.IdRequest{
		Id: "1",
	})
	if err != nil {
		return ""
	}
	return user.Id
}
