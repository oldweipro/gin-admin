package system

import (
	"context"
	"github.com/oldweipro/gin-admin/pkg/app"
	"github.com/oldweipro/gin-admin/pkg/utils"
	"go.uber.org/zap"

	"github.com/oldweipro/gin-admin/model/system"
)

type JwtService struct{}

var JwtServiceApp = new(JwtService)

//@author: [oldweipro](https://github.com/oldweipro)
//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = app.DBClient.Create(&jwtList).Error
	if err != nil {
		return
	}
	app.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := app.BlackCache.Get(jwt)
	return ok
	// err := global.DBClient.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	// isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	// return !isNotFound
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: redisJWT string, err error

func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = app.RedisClient.Get(context.Background(), userName).Result()
	return redisJWT, err
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := utils.ParseDuration(app.Config.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = app.RedisClient.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func LoadAll() {
	var data []string
	err := app.DBClient.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		app.Logger.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		app.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
