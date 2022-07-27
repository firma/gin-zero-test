package useractivity

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"miya/api/internal/model"
	"miya/api/internal/query"
	"time"
)

type (
	IUserActivityRepo interface {
		GetBy(ctx context.Context, userID int64) (*model.UserActivity, error)
		Create(ctx context.Context, ua *model.UserActivity) error
		UpdateLoginAt(ctx context.Context, userID int64) error
	}

	UserActivityRepo struct {
	}
)

func NewUserActivityRepo() IUserActivityRepo {
	return UserActivityRepo{}
}

func (r UserActivityRepo) GetBy(ctx context.Context, userID int64) (*model.UserActivity, error) {
	q := query.Ctx(ctx).UserActivity

	ent, err := q.WithContext(ctx).Where(
		q.UserID.Eq(userID),
	).Take()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return new(model.UserActivity), nil
		}

		return nil, err
	}

	return ent, nil
}

func (r UserActivityRepo) Create(ctx context.Context, ua *model.UserActivity) error {
	q := query.Ctx(ctx).UserActivity

	return q.WithContext(ctx).Create(ua)
}

func (r UserActivityRepo) UpdateLoginAt(ctx context.Context, userID int64) error {
	q := query.Ctx(ctx).UserActivity

	_, err := q.WithContext(ctx).Where(
		q.UserID.Eq(userID),
	).UpdateSimple(
		q.LatestLoginAt.Value(time.Now().Unix()),
	)

	return err

}
