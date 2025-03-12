package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

// InitRedis 初始化 Redis 客户端
func InitRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.32.137:6379",
		Password: "redis", // no password set
		DB:       0,       // use default DB
	})

	// 测试 Redis 连接
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return client, nil
}

// CloseRedis 关闭 Redis 客户端连接
func CloseRedis(redisClient *redis.Client) {
	err := redisClient.Close()
	if err != nil {
		log.Fatal(err)
	}
}
