package logic

import (
	"context"
	"gorm.io/gorm"
	"miya/api/internal/config"
	"miya/api/internal/model"
	"time"
)

type UserLogic struct {
	UserDbManager *gorm.DB
	ctx           context.Context
}

func NewUserInfoLogic(ctx context.Context) *UserLogic {
	return &UserLogic{
		ctx:           ctx,
		UserDbManager: config.Connection("user"),
	}
}

func (l *UserLogic) UserInfo() (*model.User, error) {
	// 查询用户是否存在
	user := model.User{
		Username:   "test",
		CreateTime: 0,
		UpdateTime: 0,
		ModifyTime: time.Time{},
		DeleteTime: 0,
	}
	err := l.UserDbManager.Create(&user).Error
	return &user, err
}
