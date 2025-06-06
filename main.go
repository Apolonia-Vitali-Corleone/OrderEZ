package main

//
//import (
//	"OrderEZ/internal/app/handler"
//	"OrderEZ/internal/app/service"
//	"OrderEZ/internal/infrastructure/database"
//	"OrderEZ/internal/infrastructure/messaging"
//	"OrderEZ/internal/infrastructure/server"
//	"OrderEZ/internal/util"
//	"log"
//	handler2 "order-service/internal/handler"
//	service2 "order-service/internal/service"
//	handler3 "user-service/internal/handler"
//	service3 "user-service/internal/service"
//)
//
//func main() {
//	// 初始化数据库
//	mysql, sqlDB, err := database.InitMySQL()
//	if err != nil {
//		log.Fatalf("初始化数据库失败！: %v", err)
//	}
//	defer database.CloseMySQL(sqlDB)
//
//	// 初始化 Redis
//	redis, err := database.InitRedis()
//	defer database.CloseRedis(redis)
//
//	// 初始化 RabbitMQ
//	mq, err := messaging.NewRabbitMQ()
//	defer mq.Close()
//	rabbitMQ, err := messaging.NewRabbitMQ() // 创建 RabbitMQ 实例
//	defer rabbitMQ.Close()
//
//	idGen, err := util.NewSnowflake(1, 1) // workerID=1, datacenterID=1
//	if err != nil {
//		log.Fatal("创建雪花ID生成器失败:", err)
//	}
//
//	// 初始化服务
//	userService := service3.NewUserService(mysql, redis, rabbitMQ)
//	orderService := service2.NewOrderService(mysql, redis, rabbitMQ)
//	orderItemService := service2.NewOrderItemService(mysql, redis, rabbitMQ)
//
//	goodService := service.NewGoodService(mysql, redis, rabbitMQ)
//	cartService := service.NewCartService(mysql, redis, rabbitMQ)
//	cartDetailService := service.NewCartDetailService(mysql, redis, rabbitMQ)
//	seckillGoodService := service.NewSeckillGoodService(mysql, redis, rabbitMQ)
//
//	// 初始化处理器
//	userHandler := handler3.NewUserHandler(userService)
//	orderHandler := handler2.NewOrderHandler(orderService, orderItemService, idGen)
//	goodHandler := handler.NewGoodHandler(goodService)
//	cartHandler := handler.NewCartHandler(cartService, cartDetailService)
//	seckillGoodHandler := handler.NewSeckillGoodHandler(seckillGoodService)
//
//	//// 启动服务器
//	srv := server.NewServer(userHandler, orderHandler, goodHandler, cartHandler, seckillGoodHandler)
//	srv.Run()
//}
