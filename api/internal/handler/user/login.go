package user

import (
	"context"
	"miya/common/errno"
	"miya/common/httpx"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h Handler) Login(ctx context.Context) httpx.Response {
	var req LoginReq
	if err := httpx.ShouldBindFromServerContext(ctx, &req); err != nil {
		return errno.ParamValidationErr.WithReason(err)
	}

	if err := h.userLogic.Login(ctx, req.Username, req.Password); err != nil {
		return errno.SysErr.WithReason(err)
	}

	return errno.OK
}
