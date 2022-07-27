package logic

import (
	"context"
	"fmt"

	userrepo "miya/user/internal/repository/user"
	"miya/user/internal/svc"
	"miya/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	userRepo userrepo.IUserRepo // TODO wire
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		userRepo: userrepo.NewUserRepo(),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserResp, error) {
	ent, err := l.userRepo.GetByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}

	if ent.IsEmpty() {
		return nil, fmt.Errorf("user not found") // TODO defined error
	}

	return &user.GetUserResp{
		UserId:   ent.ID,
		Username: ent.Username,
		Nickname: ent.Nickname,
	}, nil
}
