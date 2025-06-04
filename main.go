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
	defer database.CloseMySQL(sqlDB)

	// 初始化 Redis
	redis, err := database.InitRedis()
	defer database.CloseRedis(redis)

	// 初始化 RabbitMQ
	mq, err := messaging.NewRabbitMQ()
	defer mq.Close()
	rabbitMQ, err := messaging.NewRabbitMQ() // 创建 RabbitMQ 实例
	defer rabbitMQ.Close()

	// 初始化服务
	userService := service.NewUserService(mysql, redis, rabbitMQ)
	orderService := service.NewOrderService(mysql, redis, rabbitMQ)
	goodService := service.NewGoodService(mysql, redis, rabbitMQ)
	cartService := service.NewCartService(mysql, redis, rabbitMQ)
	cartDetailService := service.NewCartDetailService(mysql, redis, rabbitMQ)

	// 初始化处理器
	authHandler := handler.NewUserHandler(userService)
	orderHandler := handler.NewOrderHandler(orderService)
	goodHandler := handler.NewGoodHandler(goodService)
	cartHandler := handler.NewCartHandler(cartService, cartDetailService)

	//// 启动服务器
	srv := server.NewServer(authHandler, orderHandler, goodHandler, cartHandler)
	srv.Run()
}
