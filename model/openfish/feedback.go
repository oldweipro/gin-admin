// 自动生成模板Feedback
package openfish

import (
	"github.com/oldweipro/gin-admin/global"
)

// Feedback 结构体
type Feedback struct {
	global.GVA_MODEL
	Feedback_text string `json:"feedback_text" form:"feedback_text" gorm:"column:feedback_text;comment:反馈的文字内容;"`
	User_id       *int   `json:"user_id" form:"user_id" gorm:"column:user_id;comment:反馈用户的ID;"`
	Parent_id     *int   `json:"parent_id" form:"parent_id" gorm:"column:parent_id;comment:反馈的上级ID，主要用于回复;"`
	CreatedBy     uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy     uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy     uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Feedback 表名
func (Feedback) TableName() string {
	return "feedback"
}
