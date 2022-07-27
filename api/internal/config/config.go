package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"miya/common/stores/gormx"
)

type Config struct {
	rest.RestConf
	DB      gormx.Config
	UserRpc zrpc.RpcClientConf
}
