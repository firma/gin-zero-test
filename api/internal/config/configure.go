package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	config2 "miya/common/config"
)

type Configure struct {
	Name      string                        `json:"name" yaml:"name"`
	Host      string                        `json:"host" yaml:"host"`
	Port      string                        `json:"port" yaml:"port"`
	RpcClient map[string]zrpc.RpcClientConf `json:"RpcClient" yaml:"RpcClient"`
	Databases map[string]config2.Database   `json:"databases" yaml:"databases"`
}
