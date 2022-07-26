package middleware

import "github.com/gin-gonic/gin"

type AuthMiddleware struct {
}

func NewAuthMiddleware() AuthMiddleware {
	return AuthMiddleware{}
}

func (m AuthMiddleware) Do() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
