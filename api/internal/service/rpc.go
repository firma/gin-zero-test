package service

import (
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/zrpc"
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
	rpcConfig := config.GetConfig().UserRpcConf
	client := zrpc.MustNewClient(rpcConfig)
	rp := user.NewUser(client)
	return &RpcService{
		UserRpc:     rp,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(config.GetConfig().TransformConf)),
	}
}
