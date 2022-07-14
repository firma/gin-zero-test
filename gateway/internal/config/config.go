package config

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

var c Config

type Config struct {
	rest.RestConf
	UserRpc  zrpc.RpcClientConf
	UsersRpc zrpc.RpcClientConf
}

func GetConfig() Config {
	return c
}

func init() {
	conf.MustLoad("etc/gateway.yaml", &c)
}
