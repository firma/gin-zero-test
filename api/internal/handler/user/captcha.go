package user

import (
	"github.com/gin-gonic/gin"
)

type CaptchaReq struct {
	Username string `json:"username" binding:"required"`
}

func (h Handler) Captcha(ctx *gin.Context) {
	ctx.Data(200, "image/png", []byte("1234"))
}
