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
	authService := service.NewAuthService(redis, mysql)
	orderService := service.NewOrderService(mysql, rabbitMQ)

	// 初始化处理器
	authHandler := handler.NewAuthHandler(authService)
	orderHandler := handler.NewOrderHandler(orderService)

	//// 启动服务器
	srv := server.NewServer(authHandler, orderHandler)
	srv.Run()
}
