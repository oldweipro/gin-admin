package initialize

import (
	"bufio"
	"github.com/oldweipro/gin-admin/pkg/app"
	"github.com/oldweipro/gin-admin/pkg/utils"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"os"
	"strings"
)

func OtherInit() {
	// 解析 JWT 的过期时间配置
	// 使用 utils.ParseDuration 解析配置中的过期时间，返回时长（dr）和错误（err）
	dr, err := utils.ParseDuration(app.Config.JWT.ExpiresTime)
	if err != nil {
		// 如果解析时发生错误，程序会直接 panic 并终止
		panic(err)
	}

	// 解析 JWT 的缓冲时间配置
	// 使用 utils.ParseDuration 解析配置中的缓冲时间，忽略返回值，只处理错误（err）
	_, err = utils.ParseDuration(app.Config.JWT.BufferTime)
	if err != nil {
		// 如果解析时发生错误，程序会直接 panic 并终止
		panic(err)
	}

	// 初始化本地缓存 BlackCache
	// 使用解析出的过期时间 dr 设置缓存的默认过期时间
	app.BlackCache = local_cache.NewCache(
		// 设置默认的缓存过期时间为 dr
		local_cache.SetDefaultExpire(dr),
	)

	// 打开 go.mod 文件
	file, err := os.Open("go.mod")
	if err != nil {
		// 如果文件打开失败，直接返回
		return
	}
	// 在函数退出时关闭文件
	defer file.Close()

	// 如果 AutoCode 的 Module 还没有配置
	if app.Config.AutoCode.Module == "" {
		// 使用 bufio.Scanner 逐行扫描 go.mod 文件
		scanner := bufio.NewScanner(file)
		// 读取文件的第一行
		if scanner.Scan() {
			// 去掉第一行的 "module " 前缀，获取模块名称
			moduleName := strings.TrimPrefix(scanner.Text(), "module ")
			// 将模块名称设置到 AutoCode 的 Module 配置中
			app.Config.AutoCode.Module = moduleName
		}
	}
}
