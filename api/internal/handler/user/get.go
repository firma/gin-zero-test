package user

import (
	"context"
	"miya/api/internal/srv"
	"miya/common/errno"
	"miya/common/httpx"
	"miya/user/user"
)

type GetUserReq struct {
	Username string `form:"username" binding:"required"`
}

func (h Handler) GetUser(ctx context.Context) httpx.Response {
	var req GetUserReq
	if err := httpx.ShouldBindFromServerContext(ctx, &req); err != nil {
		return errno.ParamValidationErr.WithReason(err)
	}

	resp, err := user.NewUser(srv.Hub().User).GetUser(ctx, &user.GetUserReq{
		Username: req.Username,
	})
	if err != nil {
		return errno.SysErr.WithReason(err)
	}

	return errno.OK.WithData(GetUserResp{
		Username: resp.Username,
		Nickname: resp.Nickname,
	})
}

type GetUserResp struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}
