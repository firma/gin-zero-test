package user

import (
	"context"
	"miya/api/internal/repository/useractivity"
)

type (
	IUserLogic interface {
		Login(ctx context.Context, username string, password string) error
	}

	UserLogic struct {
		UserActivityRepo useractivity.IUserActivityRepo
	}
)

func NewUserLogic(
	userActivityRepo useractivity.IUserActivityRepo,
) IUserLogic {
	return UserLogic{
		UserActivityRepo: userActivityRepo,
	}
}
