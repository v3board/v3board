package router

import (
	"github.com/gin-gonic/gin"
)

// Router 路由实例
func Router() *gin.Engine {
	r := gin.Default()
	return r
}
