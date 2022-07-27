//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"miya/api/internal/handler"
	"miya/api/internal/logic"
	"miya/api/internal/middleware"
	"miya/api/internal/repository"
	"miya/api/internal/routes"
)

var Provider = wire.NewSet(
	handler.Provider,
	middleware.Provider,
	repository.Provider,
	logic.Provider,

	wire.Struct(new(Entry), "*"),
)

type Entry struct {
	Handlers    handler.Handlers
	Middlewares middleware.Middlewares
}

func Build() Entry {
	panic(wire.Build(Provider))
}

var entry Entry

func Setup(e Entry) {
	entry = e
}

func RegisterRoutes(r *gin.Engine) {
	routes.RegisterRoutes(r, entry.Middlewares, entry.Handlers)
}
