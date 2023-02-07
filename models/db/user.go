package db

// User 用户
type User struct {
	// 用户ID
	Id int64 `json:"id" xorm:"notnull pk autoincr 'id'"`
	// 用户uuid
	Uuid string `json:"uuid" xorm:"notnull unique 'uuid'"`
	// 权限组ID
	GroupId int `json:"group_id" xorm:"'group_id'"`
	// telegram id
	TelegramId int64 `json:"telegram_id" xorm:"'telegram_id'"`
	// 邮箱
	Email string `json:"email" xorm:"notnull 'email'"`
	// 密码
	Password string `json:"-" xorm:"notnull 'password'"`
	// 订阅token
	Token string `json:"token" xorm:"notnull unique 'token'"`
	// 是否禁用
	IsDisabled string `json:"is_disabled" xorm:"notnull default('f') 'is_disabled' comment('t-禁用,f-启用')"`
	// 禁用原因
	DisableReason string `json:"disable_reason" xorm:"disable_reason'"`
	// 是否是管理员
	IsAdmin string `json:"is_admin" xorm:"notnull 'is_admin' comment('f-不是,t-是')"`
	// 最后登录IP
	LastLoginIp string `json:"last_login_ip" xorm:"'last_login_ip'"`
	// 最后登录时间
	LastLoginAt int64 `json:"last_login_at" xorm:"'last_login_at' default(0) comment('0表示用户从未登录')"`
	// 过期时间
	ExpiredAt int64 `json:"expired_at" xorm:"notnull 'expired_at' default(0) comment('0表示永不过期')"`
	// 创建时间
	CreatedAt int64 `json:"created_at" xorm:"created notnull 'created_at'"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at" xorm:"updated notnull 'updated_at'"`
	// 删除时间
	DeletedAt int64 `json:"deleted_at" xorm:"deleted default(0) 'deleted_at'"`
}

// TableName 表名
func (*User) TableName() string {
	return "v3_user"
}
