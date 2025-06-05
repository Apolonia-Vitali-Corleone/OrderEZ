package handler

import (
	"OrderEZ/internal/app/service"
	"OrderEZ/internal/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userService.Login(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

	fmt.Println("登录成功")
}

func (h *UserHandler) Logout(c *gin.Context) {
	// 从请求头中获取令牌
	tokenStr := c.GetHeader("Authorization")
	if tokenStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少令牌"})
		return
	}
	fmt.Println("要登出的token：" + tokenStr)

	//// 去除令牌前缀 "Bearer "
	//if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
	//	tokenStr = tokenStr[7:]
	//}
	//fmt.Println(tokenStr)

	// 验证令牌
	claims, err := util.ValidateToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的令牌"})
		return
	}
	// 这里可以根据业务需求使用 claims 数据，当前登出逻辑暂不需要
	_ = claims

	// 调用服务层的登出方法
	err = h.userService.Logout(tokenStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "登出失败，请稍后再试"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "登出成功"})
	// c.JSON(http.StatusInternalServerError, gin.H{"error": "登出失败，请稍后再试"})
}

func (h *UserHandler) Register(c *gin.Context) {
	var registerData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userService.Register(registerData.Username, registerData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GetAllUsers 方法用于获取所有用户
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	// 从查询参数中获取 page 和 pageSize
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	var page, pageSize int
	var err error

	// 将 page 和 pageSize 从字符串转换为整数
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	pageSize, err = strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pageSize parameter"})
		return
	}

	users, err := h.userService.GetAllUsers(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
