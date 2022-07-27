package middleware

import "github.com/google/wire"

var Provider = wire.NewSet(
	NewAuthMiddleware,

	wire.Struct(new(Middlewares), "*"),
)

type Middlewares struct {
	Auth AuthMiddleware
}
