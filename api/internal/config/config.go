package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	util "miya/common/utils"
	"os"
	"path/filepath"
)

var Provider = wire.NewSet(NewConfig)

var c *Config
var MysqlAllClients map[string]*gorm.DB

type Config struct {
	rest.RestConf
	UserRpcConf   zrpc.RpcClientConf
	TransformConf zrpc.RpcClientConf
}

func GetConfig() *Config {
	return c
}

func NewConfig(pathInfo string) *Config {
	//conf.MustLoad(pathInfo, &c)
	//fmt.Println(c.UserRpcConf, c.TransformConf)
	InitConfig(pathInfo)

	fmt.Println(config)
	return c
}

var config = new(Configure)

func InitConfig(configFile string) {
	viper.AddConfigPath("./etc")
	viper.SetConfigType("yaml")
	viper.SetConfigName("api")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	//configFile := "etc/api.yaml"
	info, ok := util.IsExists(configFile)
	if !ok {
		if err := os.MkdirAll(filepath.Dir(configFile), 0766); err != nil {
			panic(err)
		}

		f, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if err := viper.WriteConfig(); err != nil {
			panic(err)
		}
	}
	fmt.Println(info, viper.Get("rpc_client"))
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
	GetDbInstance()
}

func GetAppConfig() Configure {
	return *config
}

func GetRpcClient(rpc string) (client zrpc.Client) {

	rpcConfig := GetAppConfig().RpcClient
	if len(rpcConfig) == 0 {
		return nil
	}

	if rpcConfig, ok := rpcConfig[rpc]; ok {
		client = zrpc.MustNewClient(rpcConfig)

	}
	return client
}
