package internal

import (
	"github.com/google/wire"
	"miya/apix/internal/app"
	"miya/apix/internal/config"
	"miya/apix/internal/rpc"
)

var ProviderSet = wire.NewSet(config.Provider, rpc.NewServiceContext, app.NewApp)
