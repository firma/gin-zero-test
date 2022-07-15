package internal

import (
	"github.com/google/wire"
	"miya/api/internal/app"
	"miya/api/internal/config"
	"miya/api/internal/service"
)

var ProviderSet = wire.NewSet(config.ConfigProvider, app.NewApp, service.NewServiceContext)
