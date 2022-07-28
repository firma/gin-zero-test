package handler

import (
	"miya/api/internal/handler/system"
	"miya/api/internal/handler/user"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	user.NewHandler,
	system.NewHandler,

	wire.Struct(new(Handlers), "*"),
)

type Handlers struct {
	User   user.Handler
	System system.Handler
}
