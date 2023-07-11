package middleware

import (
	"github.com/oldweipro/gin-admin/model/system"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/service"
	"github.com/oldweipro/gin-admin/utils"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.ConfigServer.System.Env != "develop" {
			waitUse, _ := utils.GetClaims(c)
			//获取请求的PATH
			path := c.Request.URL.Path
			obj := strings.TrimPrefix(path, global.ConfigServer.System.RouterPrefix)
			// 获取请求方法
			act := c.Request.Method
			// 获取用户的角色
			var user system.SysUser
			err := global.DB.Where("username = ?", waitUse.Username).Preload("Authorities").Preload("Authority").First(&user).Error
			if err != nil {
				c.Abort()
				return
			}
			waitUse.AuthorityId = user.AuthorityId
			sub := strconv.Itoa(int(waitUse.AuthorityId))
			e := casbinService.Casbin() // 判断策略中是否存在
			success, _ := e.Enforce(sub, obj, act)
			if !success {
				response.FailWithDetailed(gin.H{}, "权限不足", c)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
