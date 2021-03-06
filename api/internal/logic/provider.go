package logic

import (
	"github.com/google/wire"
	"miya/api/internal/logic/system"
	"miya/api/internal/logic/user"
)

var Provider = wire.NewSet(
	user.NewUserLogic,
	system.NewSystemLogic,
)
