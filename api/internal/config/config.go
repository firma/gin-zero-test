package config

import (
	"fmt"
	"github.com/google/wire"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

var ConfigProvider = wire.NewSet(NewConfig)

var c *Config

type Config struct {
	rest.RestConf
	UserRpcConf   zrpc.RpcClientConf
	TransformConf zrpc.RpcClientConf
}

func GetConfig() *Config {
	return c
}

func NewConfig() *Config {
	conf.MustLoad("etc/api.yaml", &c)
	fmt.Println(c.UserRpcConf, c.TransformConf)
	return c
}
