package system

import (
	"context"
	"miya/common/errno"
	"miya/common/httpx"
)

type CashBinReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h Handler) CasBin(ctx context.Context) httpx.Response {
	//var req user.LoginReq
	//if err := httpx.ShouldBindFromServerContext(ctx, &req); err != nil {
	//	return errno.ParamValidationErr.WithReason(err)
	//}

	if err := h.systemLogic.EnforcerTool(ctx); err != nil {
		return errno.SysErr.WithReason(err)
	}

	return errno.OK
}
