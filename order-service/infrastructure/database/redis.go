package database

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

const (
	defaultRedisAddr = "127.0.0.1:6379"
)

func redisOptions() (*redis.Options, error) {
	opts := &redis.Options{
		Addr:     defaultRedisAddr,
		Password: "",
		DB:       0,
	}

	if addr := os.Getenv("REDIS_ADDR"); addr != "" {
		opts.Addr = addr
	}

	if password, ok := os.LookupEnv("REDIS_PASSWORD"); ok {
		opts.Password = password
	}

	if dbStr := os.Getenv("REDIS_DB"); dbStr != "" {
		db, err := strconv.Atoi(dbStr)
		if err != nil {
			return nil, fmt.Errorf("invalid REDIS_DB value %q: %w", dbStr, err)
		}
		opts.DB = db
	}

	return opts, nil
}

// InitRedis 初始化 Redis 客户端
func InitRedis() (*redis.Client, error) {
	opts, err := redisOptions()
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)

	// 测试 Redis 连接
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return client, nil
}

// CloseRedis 关闭 Redis 客户端连接
func CloseRedis(redisClient *redis.Client) error {
	if redisClient == nil {
		return nil
	}

	if err := redisClient.Close(); err != nil {
		return fmt.Errorf("failed to close Redis client: %w", err)
	}
	return nil
}
