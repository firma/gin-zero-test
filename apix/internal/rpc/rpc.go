package rpc

import (
	"miya/apix/internal/config"
	"miya/transform/rpc/transformer"
	"miya/user/user"

	"github.com/gin-gonic/gin"
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
