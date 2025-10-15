//go:build integration
// +build integration

package server

import (
	"OrderEZ/internal/app/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	handler2 "order-service/internal/handler"
	"time"
	handler3 "user-service/internal/handler"
)

type Server struct {
	userHandler        *handler3.UserHandler
	orderHandler       *handler2.OrderHandler
	itemHandler        *handler.ItemHandler
	cartHandler        *handler.CartHandler
	seckillGoodHandler *handler.SeckillGoodHandler
}

func NewServer(userHandler *handler3.UserHandler, orderHandler *handler2.OrderHandler, itemHandler *handler.ItemHandler, cartHandler *handler.CartHandler, seckillGoodHandler *handler.SeckillGoodHandler) *Server {
	return &Server{
		userHandler:        userHandler,
		orderHandler:       orderHandler,
		itemHandler:        itemHandler,
		cartHandler:        cartHandler,
		seckillGoodHandler: seckillGoodHandler,
	}
}

func (s *Server) Run() {

	// 设置 gin 运行模式为 "Release" 模式，提升性能
	gin.SetMode(gin.ReleaseMode)

	// 创建一个新的 gin 引擎实例
	g := gin.New()

	// 使用 gin.Recovery()，防止程序因 panic 崩溃
	g.Use(gin.Recovery())

	// 这个一定要放在路由前
	g.Use(func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.AbortWithStatus(204)
			return
		}
	})

	// 设置 CORS 规则
	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173", "http://10.7.205.88:5173"}, // 允许的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // 允许携带 Cookie
		MaxAge:           12 * time.Hour, // 预检请求缓存时间
	}))

	//// 用户路由
	//userGroup := g.Group("/user")
	//{
	//	// 登录
	//	userGroup.POST("/login", s.userHandler.Login)
	//
	//	// 登出
	//	userGroup.POST("/logout", s.userHandler.Logout)
	//
	//	// 注册
	//	userGroup.POST("/register", s.userHandler.Register)
	//
	//	// 获取所有的用户
	//	userGroup.GET("/", s.userHandler.GetAllUsers)
	//}

	// 订单路由
	orderGroup := g.Group("/order")
	{
		// 创建订单
		orderGroup.POST("/", s.orderHandler.CreateOrder)
	}

	//// 商品路由
	//goodGroup := g.Group("/good")
	//{
	//	goodGroup.GET("/goods", s.itemHandler.GetAllGoods)
	//	goodGroup.POST("/", s.itemHandler.AddGood)
	//}

	//cartGroup := g.Group("/cart")
	//{
	//	cartGroup.GET("/", s.cartHandler.GetCart)
	//}

	//seckillGoodGroup := g.Group("/seckill_good")
	//{
	//	seckillGoodGroup.GET("", s.seckillGoodHandler.GetAllSeckillGoods)
	//}

	err := g.Run("127.0.0.1:4444")

	if err != nil {
		return
	}
}
