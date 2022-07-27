package user

import (
	"context"
	"miya/api/internal/srv"
	"miya/common/errno"
	"miya/common/httpx"
	"miya/user/user"
)

func (h Handler) TestRpc(ctx context.Context) httpx.Response {
	var req GetUserReq
	if err := httpx.ShouldBindFromServerContext(ctx, &req); err != nil {
		return errno.ParamValidationErr.WithReason(err)
	}

	resp, err := user.NewUser(srv.Hub().User).RegisterReq(ctx, &user.GetUserReq{
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
