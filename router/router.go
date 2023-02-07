package router

import (
	v1 "v3board/api/v1"

	"github.com/gin-gonic/gin"
)

// Router 路由实例
func Router() *gin.Engine {
	r := gin.Default()
	apiV1 := r.Group("/api/v1")
	{
		// 节点接口
		server := apiV1.Group("/server/UniProxy")
		{
			server.GET("/user", v1.Server.User)
			server.GET("/config", v1.Server.Config)
			server.POST("/push", v1.Server.Push)
		}
	}
	return r
}
