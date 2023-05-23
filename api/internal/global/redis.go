package global

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"inception/api/config"
)

func ParseRedis(m *config.Redis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", m.Host, m.Port),
		Password: m.Password, // no password set
		DB:       m.DB,       // use default DB
		PoolSize: 10,
	})

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	Log.Info("redis connect ping response:", ping)
	return client, nil
}
