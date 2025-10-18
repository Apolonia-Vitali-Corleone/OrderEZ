package database

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

const defaultRedisAddr = "192.168.233.136:6379"

func redisOptions() (*redis.Options, error) {
	opts := &redis.Options{
		Addr:     defaultRedisAddr,
		Password: "",
		DB:       0,
	}

	// 地址
	if addr := os.Getenv("REDIS_ADDR"); addr != "" {
		opts.Addr = addr
	}

	// 用户名/密码（如你们启用 ACL 再填）
	if user := os.Getenv("REDIS_USERNAME"); user != "" {
		opts.Username = user
	}
	if pw, ok := os.LookupEnv("REDIS_PASSWORD"); ok && pw != "" {
		// 只有非空才设置，避免误发 AUTH ""
		opts.Password = pw
	}

	// DB
	if dbStr := os.Getenv("REDIS_DB"); dbStr != "" {
		db, err := strconv.Atoi(dbStr)
		if err != nil {
			return nil, fmt.Errorf("invalid REDIS_DB value %q: %w", dbStr, err)
		}
		opts.DB = db
	}

	// ===== 关键：开启 TLS（Serverless 默认需要）=====
	// 允许用 REDIS_TLS=false 显式关闭（兼容非 TLS 环境）
	if os.Getenv("REDIS_TLS") != "false" {
		host, _, err := net.SplitHostPort(opts.Addr)
		if err != nil {
			// 如果没带端口，SplitHostPort 会报错；这种情况下直接用原字符串当 SNI
			host = opts.Addr
		}
		opts.TLSConfig = &tls.Config{
			ServerName: host, // SNI，必须与主机名一致
		}
	}

	return opts, nil
}

func InitRedis() (*redis.Client, error) {
	opts, err := redisOptions()
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	return client, nil
}

func CloseRedis(c *redis.Client) error {
	if c == nil {
		return nil
	}
	if err := c.Close(); err != nil {
		return fmt.Errorf("failed to close Redis client: %w", err)
	}
	return nil
}
