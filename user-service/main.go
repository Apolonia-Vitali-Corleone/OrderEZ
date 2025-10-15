package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"user-service/infrastructure/database"
	"user-service/infrastructure/messaging"
	"user-service/internal/handler"
	"user-service/internal/service"
	"user-service/util"
)

func main() {
	// 设置 gin 运行模式为 "Release" 模式，提升性能
	gin.SetMode(gin.ReleaseMode)

	// 创建一个新的 gin 引擎实例
	router := gin.New()

	// 使用 gin.Recovery()，防止程序因 panic 崩溃
	router.Use(gin.Recovery())

	// 自定义 CORS 中间件处理 OPTIONS 请求
	router.Use(func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// 设置 CORS 规则
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173", "http://10.7.205.88:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 初始化数据库
	mysql, sqlDB, err := database.InitMySQL()
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
	defer func() {
		if err := database.CloseMySQL(sqlDB); err != nil {
			log.Printf("关闭数据库连接失败: %v", err)
		}
	}()

	// 初始化 Redis
	redis, err := database.InitRedis()
	if err != nil {
		log.Fatalf("初始化Redis失败: %v", err)
	}
	defer func() {
		if err := database.CloseRedis(redis); err != nil {
			log.Printf("关闭Redis连接失败: %v", err)
		}
	}()

	// 初始化 RabbitMQ
	rabbitMQ, err := messaging.NewRabbitMQ()
	if err != nil {
		log.Fatalf("初始化RabbitMQ失败: %v", err)
	}
	defer func() {
		if err := rabbitMQ.Close(); err != nil {
			log.Printf("关闭RabbitMQ失败: %v", err)
		}
	}()

	// 初始化雪花ID生成器
	idGen, err := util.NewSnowflake(1, 1)
	if err != nil {
		log.Fatalf("创建雪花ID生成器失败: %v", err)
	}

	// 初始化服务层
	userService := service.NewUserService(mysql, redis, rabbitMQ)

	// 初始化处理器层
	userHandler := handler.NewUserHandler(userService, idGen)

	// 设置路由
	userGroup := router.Group("/user")
	{
		// 登录
		userGroup.POST("/login", userHandler.Login)

		// 登出
		userGroup.POST("/logout", userHandler.Logout)

		// 注册
		userGroup.POST("/register", userHandler.Register)

		// 获取所有的用户
		userGroup.GET("/", userHandler.GetAllUsers)
	}

	// 启动服务器
	log.Println("用户服务启动在 http://127.0.0.1:48482")
	if err := router.Run("127.0.0.1:48482"); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
