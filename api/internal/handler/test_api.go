package handler

import (
	"github.com/gin-gonic/gin"
	httpresponse "miya/common/basehttp"
	"net/http"
)

func Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, func() interface{} {
		return httpresponse.Success()
	}())
}
func Testtwo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, func() interface{} {
		return httpresponse.Success()
	}())
}
