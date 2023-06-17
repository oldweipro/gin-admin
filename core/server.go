package core

import (
	"fmt"
	"time"

	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/initialize"
	"github.com/oldweipro/gin-admin/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.ConfigServer.System.UseMultipoint || global.ConfigServer.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.ConfigServer.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.Logger.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 gin-admin
	当前版本:v1.0.0
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
`, address)
	global.Logger.Error(s.ListenAndServe().Error())
}
