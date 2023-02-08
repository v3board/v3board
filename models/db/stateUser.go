package db

// StateUser 用户数据统计
type StateUser struct {
	// ID
	Id int64 `json:"id" xorm:"notnull pk autoincr 'id'"`
	// 用户Id
	UserId int64 `json:"user_id" xorm:"notnull 'user_id'"`
	// 上传流量
	Up int64 `json:"up" xorm:"notnull 'up'"`
	// 上传流量
	Down int64 `json:"down" xorm:"notnull 'down'"`
	// 记录类型
	RecordType string `json:"record_type" xorm:"notnull 'record_type' comment('d-天,m-月')"`
	// 记录时间
	RecordAt int64 `json:"record_at" xorm:"notnull 'record_at'"`
	// 创建时间
	CreatedAt int64 `json:"created_at" xorm:"created notnull 'created_at'"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at" xorm:"updated notnull 'updated_at'"`
}

// TableName 表名
func (*StateUser) TableName() string {
	return "v3_state_user"
}
