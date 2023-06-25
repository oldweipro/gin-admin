// Package openfish 自动生成模板SecretKey
package openfish

import (
	"github.com/oldweipro/gin-admin/global"
	"time"
)

// OpenaiSk 结构体
type OpenaiSk struct {
	global.GVA_MODEL
	Sk     string     `json:"sk" form:"sk" gorm:"column:sk;comment:密钥;"`
	Expire *time.Time `json:"expire" form:"expire" gorm:"column:expire;comment:过期时间;"`
	Amount *int       `json:"amount" form:"amount" gorm:"column:amount;comment:余额，使用额度;"`
}

// TableName SecretKey 表名
func (OpenaiSk) TableName() string {
	return "openai_ssk"
}
