package system

import (
	"github.com/oldweipro/gin-admin/model/system"
	"github.com/oldweipro/gin-admin/pkg/app"
	"github.com/oldweipro/gin-admin/pkg/config"
	utils2 "github.com/oldweipro/gin-admin/pkg/utils"
	"go.uber.org/zap"
)

//@author: [oldweipro](https://github.com/oldweipro)
//@function: GetSystemConfig
//@description: 读取配置文件
//@return: conf config.Server, err error

type SystemConfigService struct{}

var SystemConfigServiceApp = new(SystemConfigService)

func (systemConfigService *SystemConfigService) GetSystemConfig() (conf config.Server, err error) {
	return app.Config, nil
}

// @description   set system config,
//@author: [oldweipro](https://github.com/oldweipro)
//@function: SetSystemConfig
//@description: 设置配置文件
//@param: system model.System
//@return: err error

func (systemConfigService *SystemConfigService) SetSystemConfig(system system.System) (err error) {
	cs := utils2.StructToMap(system.Config)
	for k, v := range cs {
		app.Viper.Set(k, v)
	}
	err = app.Viper.WriteConfig()
	return err
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: GetServerInfo
//@description: 获取服务器信息
//@return: server *utils.Server, err error

func (systemConfigService *SystemConfigService) GetServerInfo() (server *utils2.Server, err error) {
	var s utils2.Server
	s.Os = utils2.InitOS()
	if s.Cpu, err = utils2.InitCPU(); err != nil {
		app.Logger.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Ram, err = utils2.InitRAM(); err != nil {
		app.Logger.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils2.InitDisk(); err != nil {
		app.Logger.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}
