package system

import (
	"github.com/oldweipro/gin-admin/pkg/app"
)

type JwtBlacklist struct {
	app.BaseModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
