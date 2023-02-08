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
	if err := engine.Sync(new(db.User), new(db.Group), new(db.NodeTrojan), new(db.NodeV2ray), new(db.Setting), new(db.StateNode), new(db.StateUser)); err != nil {
		return err
	}

	// // 插入设置项
	// settings := []*db.Setting{
	// 	{
	// 		Key:         "logo",
	// 		Value:       "",
	// 		Description: "站点Logo",
	// 	},
	// 	{
	// 		Key:         "https",
	// 		Value:       "f",
	// 		Description: "强制https",
	// 	},
	// 	{
	// 		Key:         "name",
	// 		Value:       "v3board",
	// 		Description: "站点名称",
	// 	},
	// 	{
	// 		Key:         "description",
	// 		Value:       "你的个人面板",
	// 		Description: "站点描述",
	// 	},
	// 	{
	// 		Key:         "url",
	// 		Value:       "",
	// 		Description: "站点网址",
	// 	},
	// }
	// if ok, _ := engine.IsTableEmpty(new(db.Setting)); !ok {
	// 	engine.Insert(&settings)
	// }

	global.DB = engine
	return nil
}
