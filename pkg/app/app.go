package app

import (
	"fmt"
	"github.com/oldweipro/gin-admin/pkg/config"
	"github.com/oldweipro/gin-admin/pkg/utils/timer"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"

	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DBClient           *gorm.DB
	DbList             map[string]*gorm.DB
	RedisClient        redis.UniversalClient
	RedisList          map[string]redis.UniversalClient
	MongoClient        *qmgo.QmgoClient
	Config             config.Server
	Viper              *viper.Viper
	Logger             *zap.Logger
	TimerTask          = timer.NewTimerTask()
	ConcurrencyControl = &singleflight.Group{}
	Routers            gin.RoutesInfo
	ActiveDBName       *string
	BlackCache         local_cache.Cache
	lock               sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return DbList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := DbList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}

func GetRedis(name string) redis.UniversalClient {
	redisClient, ok := RedisList[name]
	if !ok || redisClient == nil {
		panic(fmt.Sprintf("redis `%s` no init", name))
	}
	return redisClient
}
