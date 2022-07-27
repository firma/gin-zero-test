package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"miya/apix/internal/config"
	"miya/apix/internal/model"
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

func (l *UserLogic) Register() (*model.SysUser, error) {
	// 查询用户是否存在
	user := model.SysUser{
		UserName:   "test",
		CreateTime: 0,
		UpdateTime: 0,
		ModifyTime: time.Time{},
		DeleteTime: 0,
	}
	err := l.UserDbManager.Create(&user).Error
	return &user, err
}

func (l *UserLogic) GetId(id int64) (*model.SysUser, error) {
	//@todo
	return nil, errors.New("test")
}
