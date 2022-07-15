package app

import (
	"miya/api/internal/service"
)

type Application struct {
	rpcService    *service.RpcService
	isCommandMode bool
}

func NewApp(rpcService *service.RpcService) (*Application, error) {
	app := &Application{
		rpcService: rpcService,
	}
	return app, nil
}
