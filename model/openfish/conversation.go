// 自动生成模板Conversation
package openfish

import (
	"github.com/oldweipro/gin-admin/global"
)

// Conversation 结构体
type Conversation struct {
	global.GVA_MODEL
	ConversationName string `json:"conversationName" form:"conversationName" gorm:"column:conversation_name;comment:会话名称;"`
	ConversationType uint   `json:"conversationType" form:"conversationType" gorm:"default:0;not null;column:conversation_type;comment:会话类型: 0:聊天,1:画图;"`
	CreatedBy        uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy        uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy        uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Conversation 表名
func (Conversation) TableName() string {
	return "conversation"
}
