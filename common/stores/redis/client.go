package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
)

type (
	Config struct {
		Host string `json:",optional"`
		Port int64  `json:",optional"`
		Type string `json:",optional"`
		Pass string `json:",optional"`
		DB   int    `json:",optional"`
	}

	DBManager struct {
		Redis *redis.Client
	}
)

func (m DBManager) GetDB(ctx context.Context) (*redis.Client, error) {
	return m.Redis.WithContext(ctx), nil
}

func MustBuildRedisDB(conf Config) *DBManager {
	m, err := BuildDBManager(conf)
	if err != nil {
		logx.ErrorStack(err)
		panic(err)
	}
	return m
}

func BuildDBManager(conf Config) (*DBManager, error) { //, logger logger.Interface

	cache := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Pass,
		DB:       conf.DB,
	})

	return &DBManager{
		Redis: cache,
	}, nil
}
