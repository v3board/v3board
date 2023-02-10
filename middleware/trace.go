package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Trace 为会话附加requestId
func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("requestId", uuid.NewString())

		ctx.Next()
	}
}
