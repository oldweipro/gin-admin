package response

import (
	"github.com/oldweipro/gin-admin/model/system"
)

type SysUserResponse struct {
	User        system.SysUser      `json:"user"`
	Permissions []map[string]string `json:"permissions"` // 用户权限
}

type LoginResponse struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
