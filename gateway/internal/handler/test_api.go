package handler

import (
	"github.com/gin-gonic/gin"
	httpresponse "miya/common/basehttp"
	"miya/gateway/internal/service"
	"net/http"
)

func Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, func() interface{} {
		c := service.NewServiceContext(ctx)
		s := c.GetUser()
		return httpresponse.Success(s)
	}())
}
