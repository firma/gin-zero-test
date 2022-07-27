package user

import (
	"miya/api/internal/handler"
	"miya/api/internal/middleware"
	"miya/common/httpx"

	"github.com/gin-gonic/gin"
)

func RegisterUser(r *gin.RouterGroup, middlewares middleware.Middlewares, handlers handler.Handlers) {
	r.GET("", httpx.Json(handlers.User.GetUser)).
		POST("login", httpx.Json(handlers.User.Login))
}
