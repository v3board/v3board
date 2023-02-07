package db

// Group 权限组
type Group struct {
	// 组ID
	Id int64 `json:"id" xorm:"notnull pk autoincr 'id'"`
	// 组名称
	Name string `json:"name" xorm:"notnull 'name'"`
	// 创建时间
	CreatedAt int64 `json:"created_at" xorm:"created notnull 'created_at'"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at" xorm:"updated notnull 'updated_at'"`
	// 删除时间
	DeletedAt int64 `json:"deleted_at" xorm:"deleted default(0) 'deleted_at'"`
}

// TableName 表名
func (*Group) TableName() string {
	return "v3_group"
}
