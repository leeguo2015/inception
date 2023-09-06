package global

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"inception/api/config"
	"time"
)

var (
	HeartBeat = 30
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
	//go heartbeatRedis(context.Background(), client)

	Log.Info("redis connect ping response:", ping)
	return client, nil
}

func heartbeatRedis(ctx context.Context, client *redis.Client) {
	timer := time.NewTimer(time.Duration(HeartBeat) * time.Second)
	select {
	case <-timer.C:
		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			Log.Info("redis心跳异常", err.Error())
		}
	case <-ctx.Done():
		timer.Stop()
		Log.Info("redis心跳退出", ctx.Err())
	}
}
