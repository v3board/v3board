package models

// ServerUniversalQueryReq 节点访问server相关接口通用query参数
type ServerUniversalQueryReq struct {
	// 节点ID
	NodeId int `form:"node_id" binding:"required"`
	// 节点类型
	NodeType string `form:"node_type" binding:"oneof=v2ray trojan"`
	// 认证token
	Token string `form:"token" binding:"required"`
}

type ServerUserRes struct {
	Users []struct {
		Id         int    `json:"id"`
		Uuid       string `json:"uuid"`
		SpeedLimit int    `json:"speed_limit"`
	} `json:"users"`
}
