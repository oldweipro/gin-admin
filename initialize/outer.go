package initialize

import (
	"github.com/songzhibin97/gkit/cache/local_cache"

	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/utils"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.ConfigServer.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.ConfigServer.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)

	global.Cache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
}
