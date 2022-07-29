package routes

import (
	"github.com/gin-gonic/gin"
	"miya/api/internal/handler"
	"miya/api/internal/middleware"
	"miya/api/internal/routes/system"
	"miya/api/internal/routes/user"
)

func RegisterRoutes(r *gin.Engine, middlewares middleware.Middlewares, handlers handler.Handlers) {
	rg := r.Group("")

	user.RegisterUser(rg.Group("api/user"), middlewares, handlers)

	system.RegisterSystem(rg.Group("api/system"), middlewares, handlers)
}
