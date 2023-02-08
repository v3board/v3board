package v1

import "github.com/gin-gonic/gin"

type settingApi struct{}

var Setting = settingApi{}

// Update 更新设置项
func (*settingApi) Update(c *gin.Context) {}
