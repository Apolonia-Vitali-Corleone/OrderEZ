package server

import (
	"OrderEZ/internal/app/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type Server struct {
	userHandler  *handler.UserHandler
	orderHandler *handler.OrderHandler
}

func NewServer(userHandler *handler.UserHandler, orderHandler *handler.OrderHandler) *Server {
	return &Server{
		userHandler:  userHandler,
		orderHandler: orderHandler,
	}
}

func (s *Server) Run() {

	// 设置 gin 运行模式为 "Release" 模式，提升性能
	gin.SetMode(gin.ReleaseMode)

	// 创建一个新的 gin 引擎实例
	g := gin.New()

	// 使用 gin.Recovery()，防止程序因 panic 崩溃
	g.Use(gin.Recovery())

	//// 设置模板函数
	//g.SetFuncMap(helperFuncs)
	//
	//// 加载 HTML 模板
	//g.LoadHTMLGlob(filepath.Join("", "./view/**/*.html"))
	//
	//// 设置静态资源路径
	//g.Static("/static", filepath.Join("", "./static"))
	//g.Static("/plugs", filepath.Join("", "./static/plugs"))
	//g.Static("/api", filepath.Join("", "./api"))

	// 设置 CORS 规则
	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://yourfrontend.com"}, // 允许的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // 允许携带 Cookie
		MaxAge:           12 * time.Hour, // 预检请求缓存时间
	}))

	// 认证路由
	authGroup := g.Group("/user")
	{
		authGroup.POST("/login", s.userHandler.Login)
		authGroup.POST("/register", s.userHandler.Register)
	}

	// 订单路由
	orderGroup := g.Group("/orders")
	{
		orderGroup.POST("", s.orderHandler.CreateOrder)
		// 其他订单路由...
	}

	err := g.Run("127.0.0.1:4444")
	if err != nil {
		return
	}
}

//var helperFuncs = template.FuncMap{
//	"jsExists": func(fpath string) bool {
//		jspath := fmt.Sprintf("./static/js/%s.js", fpath)
//		if _, err := PathExists(jspath); err == nil {
//			return true
//		}
//		return false
//	},
//	"timeFormat": TimeFormat,
//}
//
//func PathExists(path string) (bool, error) {
//	_, err := os.Stat(path)
//	if err == nil {
//		return true, nil
//	}
//	if os.IsNotExist(err) {
//		return false, nil
//	}
//	return false, err
//}
//
//func TimeFormat(sec int64) string {
//	if sec < 1 {
//		return ""
//	}
//	return time.Unix(sec, 0).Format("2006-01-02")
//}
