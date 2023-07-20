// 自动生成模板Feedback
package openfish

import (
	"github.com/oldweipro/gin-admin/global"
)

// Feedback 结构体
type Feedback struct {
	global.Model
	FeedbackText string `json:"feedbackText" form:"feedbackText" gorm:"column:feedback_text;type:longtext;comment:反馈的文字内容;"`
	ParentId     *int   `json:"parentId" form:"parentId" gorm:"default:0;not null;column:parent_id;comment:反馈的上级ID，主要用于回复;"`
	CreatedBy    uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint   `gorm:"column:deleted_by;comment:删除者"`
}

// FeedbackVo 结构体
type FeedbackVo struct {
	Feedback
	ReplyText string `json:"replyText" form:"replyText"`
}

// TableName Feedback 表名
func (Feedback) TableName() string {
	return "feedback"
}
