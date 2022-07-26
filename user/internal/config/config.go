package config

import (
	"miya/common/stores/gormx"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB gormx.Config
}
