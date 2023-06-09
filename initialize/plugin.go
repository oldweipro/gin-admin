package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/middleware"
	"github.com/oldweipro/gin-admin/plugin/email"
	"github.com/oldweipro/gin-admin/utils/plugin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}

func InstallPlugin(Router *gin.Engine) {
	PublicGroup := Router.Group("")
	fmt.Println("无鉴权插件安装==》", PublicGroup)
	PrivateGroup := Router.Group("")
	fmt.Println("鉴权插件安装==》", PrivateGroup)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//  添加跟角色挂钩权限的插件 示例 本地示例模式于在线仓库模式注意上方的import 可以自行切换 效果相同
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.ConfigServer.Email.To,
		global.ConfigServer.Email.From,
		global.ConfigServer.Email.Host,
		global.ConfigServer.Email.Secret,
		global.ConfigServer.Email.Nickname,
		global.ConfigServer.Email.Port,
		global.ConfigServer.Email.IsSSL,
	))
}
