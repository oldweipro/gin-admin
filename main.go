package main

import (
	"github.com/oldweipro/gin-admin/initialize"
	"github.com/oldweipro/gin-admin/pkg/app"
	"github.com/oldweipro/gin-admin/pkg/core"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Gin-Admin-App Swagger API接口文档
// @version                     v1.0.0
// @description                 轻量级go gin后台权限管理系统
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	app.Viper = core.Viper() // 初始化Viper
	initialize.OtherInit()   // JWT本地缓存和获取模块名称
	app.Logger = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(app.Logger)
	app.DBClient = initialize.NewGorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if app.DBClient != nil {
		//initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := app.DBClient.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
