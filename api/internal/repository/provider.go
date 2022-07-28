package repository

import (
	"github.com/google/wire"
	"miya/api/internal/repository/system"
	"miya/api/internal/repository/useractivity"
)

var Provider = wire.NewSet(
	useractivity.NewUserActivityRepo,
	system.NewCasBinRepo,
)
