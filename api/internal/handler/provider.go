package handler

import (
	"miya/api/internal/handler/user"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	user.NewHandler,

	wire.Struct(new(Handlers), "*"),
)

type Handlers struct {
	User user.Handler
}
