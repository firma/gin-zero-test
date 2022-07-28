package system

import (
	"context"
	"github.com/casbin/casbin/v2"
	"miya/api/internal/repository/system"
)

type (
	ISystemLogic interface {
		EnforcerTool(ctx context.Context) *casbin.Enforcer
	}
	SystemLogic struct {
		CasBinRepo system.ICasBinRepo
	}
)

func NewSystemLogic(
	casBinRepo system.ICasBinRepo,
) ISystemLogic {
	return SystemLogic{
		CasBinRepo: casBinRepo,
	}
}
