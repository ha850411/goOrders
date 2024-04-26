package service

import (
	"context"
	"fmt"
	"goOrders/conf"

	"github.com/go-redis/redis/v8"
)

func GetRedisClient() (*redis.Client, string) {

	var errorMsg string

	// 创建Redis客户端连接
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Settings.Redis.Host, conf.Settings.Redis.Port),
		Password: conf.Settings.Redis.Password,
		DB:       0,
	})

	// 测试Redis连接
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		errorMsg = "redis connect failed: " + err.Error()
	}

	return client, errorMsg
}
