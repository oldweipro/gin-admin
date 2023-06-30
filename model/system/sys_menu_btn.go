package system

import "github.com/oldweipro/gin-admin/global"

type SysBaseMenuBtn struct {
	global.Model
	Name          string `json:"name" gorm:"comment:按钮关键key"`
	Desc          string `json:"desc" gorm:"按钮备注"`
	SysBaseMenuID uint   `json:"sysBaseMenuID" gorm:"comment:菜单ID"`
}
