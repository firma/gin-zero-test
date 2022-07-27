package logic

import (
	"context"
	"miya/user/internal/model"
	userrepo "miya/user/internal/repository/user"
	"miya/user/internal/svc"
	"miya/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterReqLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	userRepo userrepo.IUserRepo // TODO wire
}

func NewRegisterReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterReqLogic {
	return &RegisterReqLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		userRepo: userrepo.NewUserRepo(),
	}
}

func (l *RegisterReqLogic) RegisterReq(in *user.GetUserReq) (*user.GetUserResp, error) {
	// todo: add your logic here and delete this line
	userInfo := model.User{
		Username: in.Username,
		Nickname: in.Username,
	}
	err := l.userRepo.Register(l.ctx, &userInfo)

	if err != nil {
		return nil, err
	}

	return &user.GetUserResp{
		UserId:   userInfo.ID,
		Username: userInfo.Username,
		Nickname: userInfo.Nickname,
	}, nil
}
