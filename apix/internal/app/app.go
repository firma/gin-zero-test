package app

import (
	"miya/apix/internal/rpc"
)

type Application struct {
	rpcService    *rpc.RpcService
	isCommandMode bool
}

func NewApp(rpcService *rpc.RpcService) (*Application, error) {
	app := &Application{
		rpcService: rpcService,
	}
	return app, nil
}
