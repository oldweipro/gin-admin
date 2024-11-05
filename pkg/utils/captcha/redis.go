package captcha

import (
	"context"
	"github.com/oldweipro/gin-admin/pkg/app"
	"time"

	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
		Context:    context.TODO(),
	}
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}

func (rs *RedisStore) Set(id string, value string) error {
	err := app.RedisClient.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		app.Logger.Error("RedisStoreSetError!", zap.Error(err))
		return err
	}
	return nil
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := app.RedisClient.Get(rs.Context, key).Result()
	if err != nil {
		app.Logger.Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err := app.RedisClient.Del(rs.Context, key).Err()
		if err != nil {
			app.Logger.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}
