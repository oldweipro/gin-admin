// 自动生成模板Prompt
package openfish

import (
	"github.com/oldweipro/gin-admin/global"
)

// Prompt 结构体
type Prompt struct {
	global.Model
	Name        string `json:"name" form:"name" gorm:"column:name;comment:表示提示词的名称。;"`
	Description string `json:"description" form:"description" gorm:"column:description;comment:表示提示词的详细描述信息。;"`
	UseFee      *int   `json:"useFee" form:"useFee" gorm:"column:use_fee;comment:表示提示词的价格。0表示无限制;"`
	Content     string `json:"content" form:"content" gorm:"column:content;comment:提示词内容;"`
	IsShare     *int   `json:"isShare" form:"isShare" gorm:"column:is_share;comment:是否分享;"`
	ShortcutKey string `json:"shortcutKey" form:"shortcutKey" gorm:"column:shortcut_key;comment:快捷键;"`
	Category    *int   `json:"category" form:"category" gorm:"column:category;comment:类别分类;"`
	UseDuration *int   `json:"useDuration" form:"useDuration" gorm:"column:use_duration;comment:使用时间，0表示无限制;"`
}

// TableName Prompt 表名
func (Prompt) TableName() string {
	return "prompt"
}
