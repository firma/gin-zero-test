package system

import (
	"miya/api/internal/logic/system"
)

type Handler struct {
	systemLogic system.ISystemLogic
}

func NewHandler(
	systemLogic system.ISystemLogic,
) Handler {
	return Handler{
		systemLogic: systemLogic,
	}
}
