package global

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `json:"id" form:"id" gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      `json:"createdAt" form:"createdAt"`     // 创建时间
	UpdatedAt time.Time      `json:"updatedAt" form:"updatedAt"`     // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                 // 删除时间
}
