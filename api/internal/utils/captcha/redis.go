package captcha

import (
	"context"
	"inception/api/internal/global"
	"time"

	"github.com/mojocn/base64Captcha"
)

func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
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
	err := global.REDIS.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		global.Log.Error("RedisStoreSetError!", err.Error())
	}
	return err
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := global.REDIS.Get(rs.Context, key).Result()
	if err != nil {
		global.Log.Error("RedisStoreGetError!", err.Error())
		return ""
	}
	if clear {
		err := global.REDIS.Del(rs.Context, key).Err()
		if err != nil {
			global.Log.Error("RedisStoreClearError!", err.Error())
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
