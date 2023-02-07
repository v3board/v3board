package db

// NodeV2ray v2ray节点
type NodeV2ray struct {
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
	// TLS支持
	Tls string `json:"tls" xorm:"notnull 'tls'"`
	// TLS设置
	TlsSetting string `json:"tls_setting" xorm:"'tls_setting'"`
	// 传输协议
	Network string `json:"network" xorm:"notnull text 'network'"`
	// 协议设置
	NetworkSetting string `json:"network_setting" xorm:"text 'network_setting'"`
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
func (*NodeV2ray) TableName() string {
	return "v3_node_v2ray"
}
