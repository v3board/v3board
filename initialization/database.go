package initialization

import (
	"path"
	"v3board/global"
	"v3board/models/db"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

// InitDatabase 初始化数据库链接
func InitDatabase() error {
	engine, err := xorm.NewEngine("sqlite3", path.Join(global.WorkDir, "v3board.db"))
	if err != nil {
		return err
	}

	// 迁移表结构
	if err := engine.Sync(new(db.User), new(db.Group), new(db.NodeTrojan), new(db.NodeV2ray)); err != nil {
		return err
	}

	global.DB = engine
	return nil
}
