package handler

import "github.com/gin-gonic/gin"

func RegisterGatewayRouters(router *gin.Engine) {
	router.GET("/test/rpc", TestRpc)
	router.GET("/test/config", TestConfig)
	router.GET("/test/db", Testdb)
}
