package lib

import "github.com/gin-gonic/gin"

// GetRequestId 获取会话ID
func GetRequestId(c *gin.Context) string {
	return c.GetString("requestId")
}
