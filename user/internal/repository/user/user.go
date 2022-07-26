package user

import (
	"context"
	"errors"
	"miya/user/internal/model"
	"miya/user/internal/query"

	"gorm.io/gorm"
)

type (
	IUserRepo interface {
		GetByUsername(ctx context.Context, username string) (*model.User, error)
	}

	userRepo struct{}
)

func NewUserRepo() IUserRepo {
	return userRepo{}
}

func (r userRepo) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	q := query.Ctx(ctx).User

	u, err := q.WithContext(ctx).
		Where(q.Username.Eq(username)).
		Take()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return new(model.User), nil
		}

		return nil, err
	}

	return u, nil
}
