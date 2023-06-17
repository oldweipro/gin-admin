package initialize

import (
	"context"

	"github.com/oldweipro/gin-admin/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.ConfigServer.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Logger.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.Logger.Info("redis connect ping response:", zap.String("pong", pong))
		global.RedisClient = client
	}
}
