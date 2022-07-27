package handler

import (
	"github.com/gin-gonic/gin"
	config2 "miya/apix/internal/config"
	"miya/apix/internal/logic"
	"miya/apix/internal/rpc"
	httpresponse "miya/common/basehttp"
	"miya/common/errorcode"
	user2 "miya/user/user"
	"net/http"
)

func TestRpc(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, func() interface{} {
		user, _ := rpc.NewServiceContext().UserRpc.GetUser(ctx, &user2.GetUserReq{
			Username: "1",
		})

		return httpresponse.Success(user)
	}())
}
func TestConfig(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, func() interface{} {
		config := config2.GetAppConfig()
		return httpresponse.Success(config)
	}())
}

func Testdb(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, func() interface{} {
		user, err := logic.NewUserInfoLogic(ctx).Register()
		if err != nil {
			return httpresponse.Error(errorcode.HandlerError.SetErrorMsg(err.Error()))
		}

		return httpresponse.Success(user)
	}())
}
