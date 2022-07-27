package user

import "miya/api/internal/logic/user"

type Handler struct {
	userLogic user.IUserLogic
}

func NewHandler(
	userLogic user.IUserLogic,
) Handler {
	return Handler{
		userLogic: userLogic,
	}
}
