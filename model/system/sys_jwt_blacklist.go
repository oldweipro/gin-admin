package system

import (
	"github.com/oldweipro/gin-admin/global"
)

type JwtBlacklist struct {
	global.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}
