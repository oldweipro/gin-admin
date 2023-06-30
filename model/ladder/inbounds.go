// 自动生成模板Inbounds
package ladder

import (
	"github.com/oldweipro/gin-admin/global"
)

// Inbounds 结构体
type Inbounds struct {
	global.Model
	ClientId       string `json:"clientId" form:"clientId" gorm:"column:client_id;comment:vmess协议客户端id，也就是密码;"`
	Up             *int64 `json:"up" form:"up" gorm:"column:up;comment:上行流量;size:64;"`
	Down           *int64 `json:"down" form:"down" gorm:"column:down;comment:下行流量;size:64;"`
	Total          *int64 `json:"total" form:"total" gorm:"column:total;comment:流量限制，0:无限制;size:64;"`
	Remark         string `json:"remark" form:"remark" gorm:"column:remark;comment:入站规则名称;"`
	Enable         *bool  `json:"enable" form:"enable" gorm:"column:enable;comment:是否启用;"`
	ExpiryTime     *int64 `json:"expiryTime" form:"expiryTime" gorm:"column:expiry_time;comment:是一个13位的时间戳;size:13;"`
	Port           string `json:"port" form:"port" gorm:"column:port;comment:端口;size:12;"`
	Protocol       string `json:"protocol" form:"protocol" gorm:"column:protocol;comment:协议;"`
	Settings       string `json:"settings" form:"settings" gorm:"column:settings;comment:基本信息;size:500;"`
	StreamSettings string `json:"streamSettings" form:"streamSettings" gorm:"column:stream_settings;comment:其他信息;size:500;"`
	Sniffing       string `json:"sniffing" form:"sniffing" gorm:"column:sniffing;comment:默认就行;size:500;"`
	Listen         string `json:"listen" form:"listen" gorm:"column:listen;comment:监听IP默认留空;"`
	Uid            *uint  `json:"uid" form:"uid" gorm:"colum:uid;comment:关联用户ID，也是入站规则的ID，在修改当前入站规则的时候会用到;"`
	Sid            *uint  `json:"sid" form:"sid" gorm:"colum:sid;comment:关联服务器ID;"`
	Link           string `json:"link" form:"link" gorm:"colum:link;comment:vmess链接;size:500"`
	Link64         string `json:"link64" form:"link64" gorm:"colum:link64;comment:vmess链接base64加密后的;size:500"`
	CreatedBy      uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy      uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy      uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Inbounds 表名
func (Inbounds) TableName() string {
	return "inbounds"
}
