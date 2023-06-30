// 自动生成模板ServerNode
package ladder

import (
	"github.com/oldweipro/gin-admin/global"
)

// ServerNode 结构体
type ServerNode struct {
	global.Model
	ServerName   string `json:"serverName" form:"serverName" gorm:"column:server_name;comment:服务器节点名称;"`
	ServerHost   string `json:"serverHost" form:"serverHost" gorm:"column:server_host;comment:服务器主机地址;"`
	ServerPort   *int   `json:"serverPort" form:"serverPort" gorm:"column:server_port;comment:服务器端口;"`
	ServerStatus *int   `json:"serverStatus" form:"serverStatus" gorm:"column:server_status;comment:服务器状态;"`
	Describe     string `json:"describe" form:"describe" gorm:"column:describe;comment:描述;"`
	Bandwidth    string `json:"bandwidth" form:"bandwidth" gorm:"column:bandwidth;comment:服务器带宽;"`
	Username     string `json:"username" form:"username" gorm:"colum:username;comment:用户名;"`
	Password     string `json:"password" form:"password" gorm:"colum:password;comment:密码;"`
	PemFile      string `json:"pemFile" form:"pemFile" gorm:"colum:pemFile;comment:pem文件;"`
	KeyFile      string `json:"keyFile" form:"keyFile" gorm:"colum:keyFile;comment:key文件;"`
	Region       string `json:"region" form:"region" gorm:"column:region;comment:服务器所在地;"`
	Domain       string `json:"domain" form:"domain" gorm:"column:domain;comment:域名;"`
	Cookie       string `json:"cookie" form:"cookie" gorm:"column:cookie;comment:cookie;size:500;"`
	CreatedBy    uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName ServerNode 表名
func (ServerNode) TableName() string {
	return "server_node"
}
