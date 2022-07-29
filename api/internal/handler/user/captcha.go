package user

import (
	"github.com/gin-gonic/gin"
	util "miya/common/utils"
)

type CaptchaReq struct {
	Username string `json:"username" binding:"required"`
}

func (h Handler) Captcha(ctx *gin.Context) {
	b := util.ImgText(100, 50, "1231")
	ctx.Data(200, "image/png", b)
}
