package core

import (
	"fmt"
	"github.com/oldweipro/gin-admin/initialize"
	"github.com/oldweipro/gin-admin/pkg/app"
	"github.com/oldweipro/gin-admin/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if app.Config.System.UseMultipoint || app.Config.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
		initialize.RedisList()
	}

	if app.Config.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}
	// 从db加载jwt数据
	if app.DBClient != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", app.Config.System.Addr)
	s := initServer(address, Router)

	app.Logger.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
    项目启动成功
	默认自动化文档地址: http://127.0.0.1%s/swagger/index.html
`, address)
	app.Logger.Error(s.ListenAndServe().Error())
}
