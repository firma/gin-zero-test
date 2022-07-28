package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"miya/common/stores/gormx"
	"miya/common/stores/redis"
)

type Config struct {
	rest.RestConf
	DB      gormx.Config
	Redis   redis.Config
	UserRpc zrpc.RpcClientConf
}
