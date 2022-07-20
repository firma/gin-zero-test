package logic

import (
	"miya/api/internal/service"
	"miya/transform/rpc/transformer"
	user2 "miya/user/rpc/types/user"
)

type Logic struct {
	serviceContext *service.RpcService
}

func NewLogic(serviceContext *service.RpcService) *Logic {
	return &Logic{
		serviceContext: serviceContext,
	}
}

func (l *Logic) GetTest() string {
	//@todo test
	user, err := l.serviceContext.UserRpc.GetUser(l.serviceContext.Context, &user2.IdRequest{
		Id: "1",
	})
	if err != nil {
		return ""
	}
	return user.Id
}

func (l *Logic) GetTest2() string {
	//@todo test
	tran, err := l.serviceContext.Transformer.Expand(l.serviceContext.Context, &transformer.ExpandReq{Shorten: "f35b2a"})
	if err != nil {
		return ""
	}
	if len(tran.Url) > 0 {
		// 代码测试
		user, err := l.serviceContext.UserRpc.GetUser(l.serviceContext.Context, &user2.IdRequest{
			Id: "1",
		})
		if err != nil {
			return ""
		}
		return tran.Url + user.Id
	}
	return tran.Url
}
