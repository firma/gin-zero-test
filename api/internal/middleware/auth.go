package middleware

import (
	"github.com/gin-gonic/gin"
	"miya/api/internal/logic/system"
	"miya/common/errno"
	"net/http"
)

type AuthMiddleware struct {
	systemLogic system.ISystemLogic
}

func NewAuthMiddleware(systemLogic system.ISystemLogic) AuthMiddleware {
	return AuthMiddleware{
		systemLogic: systemLogic,
	}
}

func (m AuthMiddleware) Do() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//@todo 无权操作
		s := m.systemLogic.Enforce(ctx, "admins", "operation", "/api/user/testrpc", "read")
		if s == false {
			ctx.JSON(http.StatusOK, errno.SysErr.WithReason(errno.Forbidden))
			ctx.Abort()
			return
		}
	}
}
