package middleware

import "github.com/gin-gonic/gin"

// Trace 为会话附加traceid
func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
