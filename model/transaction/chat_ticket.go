// 自动生成模板ChatTicket
package transaction

import (
	"github.com/oldweipro/gin-admin/global"
)

// ChatTicket 结构体
type ChatTicket struct {
	global.GVA_MODEL
	Amount         *int   `json:"amount" form:"amount" gorm:"column:amount;comment:数量;"`
	ExpirationTime *int   `json:"expirationTime" form:"expirationTime" gorm:"column:expiration_time;comment:过期时间，时间戳，0表示无过期时间;"`
	TicketName     string `json:"ticketName" form:"ticketName" gorm:"column:ticket_name;comment:票据名称;"`
	TicketValue    string `json:"ticketValue" form:"ticketValue" gorm:"column:ticket_value;comment:票据码;"`
	BelongTo       *int   `json:"belongTo" form:"belongTo" gorm:"column:belong_to;comment:归属;"`
	CreatedBy      uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy      uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy      uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName ChatTicket 表名
func (ChatTicket) TableName() string {
	return "chat_ticket"
}
