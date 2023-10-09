package bilibili

import "github.com/oldweipro/gin-admin/global"

type UpMember struct {
	global.Model
	Mid      uint   `json:"mid" gorm:"column:id;comment:up主ID"`
	Name     string `json:"name" gorm:"column:name;comment:昵称"`
	Sex      string `json:"sex" gorm:"column:sex;comment:性别"`
	Face     string `json:"face" gorm:"column:face;comment:头像"`
	Sign     string `json:"sign" gorm:"column:sign;type:longtext;comment:签名"`
	Level    uint   `json:"level" gorm:"column:level;comment:等级"`
	Birthday string `json:"birthday" gorm:"column:birthday;comment:生日"`
}

func (UpMember) TableName() string {
	return "up_member"
}
