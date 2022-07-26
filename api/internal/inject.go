package internal

import (
	"github.com/google/wire"
	"miya/api/internal/app"
	"miya/api/internal/config"
	"miya/api/internal/rpc"
)

var ProviderSet = wire.NewSet(config.Provider, rpc.NewServiceContext, app.NewApp)
