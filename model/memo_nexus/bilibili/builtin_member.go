package bilibili

import (
	"gorm.io/gorm"
	"time"
)

type BuiltinMember struct {
	Mid             uint           `json:"mid" gorm:"column:mid;primary_key;comment:系统账户mid"`
	Name            string         `json:"name" gorm:"column:name;comment:昵称"`
	SessData        string         `json:"sessData" gorm:"column:sess_data;type:longtext;comment:cookie"`
	BiliJct         string         `json:"biliJct" gorm:"column:bili_jct;comment:cookie"`
	DedeUserID      string         `json:"dedeUserID" gorm:"column:dede_user_id;comment:cookie"`
	Sid             string         `json:"sid" gorm:"column:sid;type:longtext;comment:cookie"`
	DedeUserIDCkMd5 string         `json:"dedeUserIDCkMd5" gorm:"column:dede_user_id_ck_md5;comment:cookie"`
	CreatedAt       time.Time      `json:"createdAt" form:"createdAt"` // 创建时间
	UpdatedAt       time.Time      `json:"updatedAt" form:"updatedAt"` // 更新时间
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`             // 删除时间
}

func (BuiltinMember) TableName() string {
	println()
	return "builtin_member"
}
