package system

import (
	"github.com/gin-gonic/gin"
	"miya/api/internal/handler"
	"miya/api/internal/middleware"
	"miya/common/httpx"
)

func RegisterSystem(r *gin.RouterGroup, middlewares middleware.Middlewares, handlers handler.Handlers) {
	r.GET("casbin", httpx.Json(handlers.System.CasBin))
}
