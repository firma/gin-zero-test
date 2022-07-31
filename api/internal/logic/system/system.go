package system

import (
	"context"
	"github.com/casbin/casbin/v2"
	"miya/api/internal/repository/system"
)

type (
	ISystemLogic interface {
		EnforcerTool(ctx context.Context) *casbin.Enforcer
		Enforce(ctx context.Context, admin, platform, path, method string) bool
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
