package logic

import (
	"context"
	"errors"
	"miya/merchants/api/internal/svc"
	"miya/merchants/api/internal/types"
	user2 "miya/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantLogic {
	return &GetMerchantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMerchantLogic) GetMerchant(req *types.MerchantReq) (resp *types.MerchantReply, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user2.IdRequest{
		Id: "1",
	})
	if err != nil {
		return nil, err
	}

	if user.Name != "test" {
		return nil, errors.New("用户不存在")
	}

	return &types.MerchantReply{
		Id:   req.Id,
		Name: "test order",
	}, nil
}
