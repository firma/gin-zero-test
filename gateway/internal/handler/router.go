package handler

import "github.com/gin-gonic/gin"

func RegisterGatewayRouters(router *gin.Engine) {
	router.GET("/test", Test)
}
