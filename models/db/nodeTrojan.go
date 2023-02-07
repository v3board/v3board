package db

// NodeTrojan trojan节点
type NodeTrojan struct {
	// 节点ID
	Id int64 `json:"id" xorm:"notnull pk autoincr 'id'"`
	// 权限组ID
	GroupId string `json:"group_id" xorm:"notnull 'group_id'"`
	// 节点名称
	Name string `json:"name" xorm:"notnull 'name'"`
	// 节点地址
	Host string `json:"host" xorm:"notnull 'host'"`
	// 节点端口
	Port int `json:"port" xorm:"notnull 'port'"`
	// 节点标签
	Tag string `json:"tag" xorm:"'tags'"`
	// 是否允许不安全
	AllowInsecure string `json:"allow_insecure" xorm:"notnull default('f') 'allow_insecure' comment('f-不允许,t-允许')"`
	// sni域名
	ServerName string `json:"server_name" xorm:"'server_name'"`
	// 是否禁用
	IsDisnabled string `json:"is_disenabled" xorm:"'is_disenabled' default('f') comment('f-启用,t-禁用')"`
	// 禁用原因
	DisenableReason string `json:"disable_reason" xorm:"'disable_reason'"`
	// 创建时间
	CreatedAt int64 `json:"created_at" xorm:"created notnull 'created_at'"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at" xorm:"updated notnull 'updated_at'"`
	// 删除时间
	DeletedAt int64 `json:"deleted_at" xorm:"deleted default(0) 'deleted_at'"`
}

// TableName 表名
func (*NodeTrojan) TableName() string {
	return "v3_node_trojan"
}
