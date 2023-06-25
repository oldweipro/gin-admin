// Package openfish 自动生成模板SecretKey
package openfish

import (
	"github.com/oldweipro/gin-admin/global"
)

// SecretKey 结构体
type SecretKey struct {
	global.GVA_MODEL
	Sk        string  `json:"sk" form:"sk" gorm:"column:sk;comment:密钥;"`
	SkName    string  `json:"skName" form:"skName" gorm:"column:sk_name;comment:密钥名字;not null;default:'sk';"`
	UserId    uint    `json:"userId" form:"userId" gorm:"column:user_id;comment:关联用户ID;"`
	Expire    *uint64 `json:"expire" form:"expire" gorm:"column:expire;comment:过期时间;not null;default:0;"`
	Amount    *int    `json:"amount" form:"amount" gorm:"column:amount;comment:余额，使用额度;not null;default:0;"`
	Status    *int    `json:"status" form:"status" gorm:"column:status;comment:状态：0封禁，1正常;not null;default:1;"`
	CreatedBy uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint    `gorm:"column:deleted_by;comment:删除者"`
}

// TableName SecretKey 表名
func (SecretKey) TableName() string {
	return "secret_key"
}
