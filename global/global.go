package global

import (
	"github.com/rs/zerolog"
	"xorm.io/xorm"
)

var (
	// 工作目录
	WorkDir string
	// 业务日志实例
	Log *zerolog.Logger
	// 数据库实例
	DB *xorm.Engine
	// 盐值，加密密码使用
	Sault string
)
