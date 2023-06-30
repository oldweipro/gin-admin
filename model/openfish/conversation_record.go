// 自动生成模板ConversationRecord
package openfish

import (
	"github.com/oldweipro/gin-admin/global"
)

// ConversationRecord 结构体
type ConversationRecord struct {
	global.Model
	ConversationId *uint  `json:"conversationId" form:"conversationId" gorm:"column:conversation_id;not null;comment:会话ID;"`
	Content        string `json:"content" form:"content" gorm:"column:content;comment:根据角色不通，user表示用户向服务端发起提问的问题,openai回复的消息，取最后完整的消息;type:longtext;"`
	Role           string `json:"role" form:"role" gorm:"column:role;comment:用户角色;"`
	CreatedBy      uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy      uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy      uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Conversation 表名
func (ConversationRecord) TableName() string {
	return "conversation_record"
}
