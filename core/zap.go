package core

import (
	"fmt"
	"github.com/oldweipro/gin-admin/core/internal"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 获取 zap.Logger
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.ConfigServer.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.ConfigServer.Zap.Director)
		_ = os.Mkdir(global.ConfigServer.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.ConfigServer.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
