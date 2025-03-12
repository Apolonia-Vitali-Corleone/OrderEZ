package server

import (
	"OrderEZ/internal/app/handler"
	"github.com/gin-gonic/gin"
)

type Server struct {
	authHandler  *handler.AuthHandler
	orderHandler *handler.OrderHandler
}

func NewServer(authHandler *handler.AuthHandler, orderHandler *handler.OrderHandler) *Server {
	return &Server{
		authHandler:  authHandler,
		orderHandler: orderHandler,
	}
}

func (s *Server) Run() {
	r := gin.Default()

	// 认证路由
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", s.authHandler.Login)
		authGroup.POST("/register", s.authHandler.Register)
	}

	// 订单路由
	orderGroup := r.Group("/orders")
	{
		orderGroup.POST("", s.orderHandler.CreateOrder)
		// 其他订单路由...
	}

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
