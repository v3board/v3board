package v1

import (
	"net/http"
	"v3board/global"
	"v3board/models"

	"github.com/gin-gonic/gin"
)

type serverApi struct{}

var Server = serverApi{}

// User 获取用户列表
func (*serverApi) User(c *gin.Context) {}

// Config 获取节点配置
func (*serverApi) Config(c *gin.Context) {}

// Push 上报节点信息
func (*serverApi) Push(c *gin.Context) {
	var req models.ServerUniversalQueryReq

	if err := c.ShouldBindQuery(&req); err != nil {
		global.Log.Err(err).Msg("node request failed")
		c.String(http.StatusInternalServerError, "query param error")
		return
	}

	// TODO 验证token

	c.String(http.StatusOK, "success")
}
