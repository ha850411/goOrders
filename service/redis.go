package service

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func GetRedisClient() (*redis.Client, string) {

	var errorMsg string

	// 创建Redis客户端连接
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // Redis服务器地址
		Password: "",               // Redis密码，如果有的话
		DB:       0,                // Redis数据库索引（默认为0）
	})

	// 测试Redis连接
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		errorMsg = "redis connect failed: " + err.Error()
	}

	return client, errorMsg
}
