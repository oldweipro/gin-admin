package main

import (
	"database/sql"
	"go.uber.org/zap"

	"github.com/oldweipro/gin-admin/core"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample ConfigServer pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.Viper = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.Logger = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.Logger)
	global.DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				global.Logger.Error("关闭数据库连接失败", zap.Error(err))
			}
		}(db)
	}
	core.RunWindowsServer()
}
