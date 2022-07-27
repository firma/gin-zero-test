package user

import (
	"context"
	"miya/api/internal/model"
	"miya/api/internal/srv"
	"miya/user/user"
	"time"
)

func (l UserLogic) Login(ctx context.Context, username string, password string) error {
	resp, err := user.NewUser(srv.Hub().User).GetUser(ctx, &user.GetUserReq{
		Username: username,
	})
	if err != nil {
		return err
	}

	// TODO auth

	ua, err := l.UserActivityRepo.GetBy(ctx, resp.UserId)
	if err != nil {
		return err
	}

	if ua.IsEmpty() {
		return l.UserActivityRepo.Create(ctx, &model.UserActivity{
			UserID:        resp.UserId,
			LatestLoginAt: time.Now().Unix(),
		})
	}

	return l.UserActivityRepo.UpdateLoginAt(ctx, resp.UserId)
}
