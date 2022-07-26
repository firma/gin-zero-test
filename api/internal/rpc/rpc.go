package rpc

import (
	"github.com/gin-gonic/gin"
	"miya/api/internal/config"
	"miya/transform/rpc/transformer"
	"miya/user/rpc/user"
)

type RpcService struct {
	UserRpc     user.User
	Transformer transformer.Transformer
	Context     *gin.Context
}

func NewServiceContext() *RpcService {
	return &RpcService{
		UserRpc: user.NewUser(config.GetRpcClient("user")),
	}
}
