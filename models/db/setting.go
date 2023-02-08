package db

// Setting 系统设置
type Setting struct {
	// ID
	Id int64 `json:"id" xorm:"notnull pk autoincr 'id'"`
	// 设置项
	Key string `json:"key" xorm:"notnull unique 'key'"`
	// 设置值
	Value string `json:"value" xorm:"'value'"`
	// 描述
	Description string `json:"description" xorm:"text 'description'"`
	// 创建时间
	CreatedAt int64 `json:"created_at" xorm:"created notnull 'created_at'"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at" xorm:"updated notnull 'updated_at'"`
}

// TableName 表名
func (*Setting) TableName() string {
	return "v3_setting"
}
