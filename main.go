package main

import (
	"OrderEZ/internal/app/handler"
	"OrderEZ/internal/app/service"
	"OrderEZ/internal/infrastructure/database"
	"OrderEZ/internal/infrastructure/messaging"
	"OrderEZ/internal/infrastructure/server"
	"log"
)

func main() {
	// 初始化数据库
	mysql, sqlDB, err := database.InitMySQL()
	if err != nil {
		log.Fatalf("初始化数据库失败！: %v", err)
	}

	// 延迟关闭数据库
	defer database.CloseMySQL(sqlDB)

	// 初始化 Redis
	redis, err := database.InitRedis()

	// 延迟关闭redis
	defer database.CloseRedis(redis)

	// 初始化 RabbitMQ
	rabbitMQ := messaging.InitRabbitMQ()

	// 初始化服务
	authService := service.NewUserService(redis, mysql)
	orderService := service.NewOrderService(mysql, rabbitMQ)

	// 初始化处理器
	authHandler := handler.NewAuthHandler(authService)
	orderHandler := handler.NewOrderHandler(orderService)

	//// 启动服务器
	srv := server.NewServer(authHandler, orderHandler)
	srv.Run()
}

//package main
//
//import (
//	"context"
//	"fmt"
//	"github.com/go-redis/redis/v8"
//	"time"
//)
//
//// RedisLock 表示 Redis 分布式锁
//type RedisLock struct {
//	client *redis.Client
//	key    string
//	value  string
//	expiry time.Duration
//}
//
//// NewRedisLock 创建一个新的 Redis 分布式锁实例
//func NewRedisLock(client *redis.Client, key, value string, expiry time.Duration) *RedisLock {
//	return &RedisLock{
//		client: client,
//		key:    key,
//		value:  value,
//		expiry: expiry,
//	}
//}
//
//// Lock 尝试获取锁
//func (l *RedisLock) Lock(ctx context.Context) (bool, error) {
//	// 使用 SET 命令并设置 NX 和 EX 参数尝试获取锁
//	set, err := l.client.SetNX(ctx, l.key, l.value, l.expiry).Result()
//	if err != nil {
//		return false, err
//	}
//	return set, nil
//}
//
//// Unlock 释放锁
//func (l *RedisLock) Unlock(ctx context.Context) (bool, error) {
//	// 定义 Lua 脚本，保证解锁操作的原子性
//	luaScript := `
//    if redis.call("GET", KEYS[1]) == ARGV[1] then
//        return redis.call("DEL", KEYS[1])
//    else
//        return 0
//    end
//    `
//	// 执行 Lua 脚本
//	result, err := l.client.Eval(ctx, luaScript, []string{l.key}, l.value).Result()
//	if err != nil {
//		return false, err
//	}
//	// 判断解锁是否成功
//	if result.(int64) == 1 {
//		return true, nil
//	}
//	return false, nil
//}
//
//func main() {
//	// 创建 Redis 客户端
//	client := redis.NewClient(&redis.Options{
//		Addr:     "192.168.32.137:6379",
//		Password: "redis",
//		DB:       0,
//	})
//
//	// 创建 Redis 分布式锁实例
//	lock := NewRedisLock(client, "distributed_lock", "unique_value", 10*time.Second)
//
//	ctx := context.Background()
//
//	// 尝试获取锁
//	locked, err := lock.Lock(ctx)
//	if err != nil {
//		fmt.Println("获取锁时出错:", err)
//		return
//	}
//
//	if locked {
//		fmt.Println("成功获取锁，开始执行临界区代码...")
//		// 模拟临界区代码执行
//		time.Sleep(5 * time.Second)
//		fmt.Println("临界区代码执行完毕，尝试释放锁...")
//
//		// 释放锁
//		unlocked, err := lock.Unlock(ctx)
//		if err != nil {
//			fmt.Println("释放锁时出错:", err)
//		} else if unlocked {
//			fmt.Println("锁已成功释放")
//		} else {
//			fmt.Println("释放锁失败，可能锁已过期或被其他客户端持有")
//		}
//	} else {
//		fmt.Println("获取锁失败，锁已被其他客户端持有")
//	}
//}
