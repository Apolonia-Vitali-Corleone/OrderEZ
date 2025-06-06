package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-service/internal/model/dto"
	"order-service/internal/service"
	"order-service/util"
)

type OrderHandler struct {
	orderService     *service.OrderService
	orderItemService *service.OrderItemService
	idGen            *util.Snowflake // ✅ 雪花 ID 生成器
}

// NewOrderHandler 构造函数需要初始化所有字段
func NewOrderHandler(orderService *service.OrderService, orderItemService *service.OrderItemService, idGen *util.Snowflake) *OrderHandler {
	return &OrderHandler{
		orderService:     orderService,
		orderItemService: orderItemService,
		idGen:            idGen,
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	// 1. 从请求头获取 Authorization 字段
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "缺少 Authorization 头"})
		return
	}

	// 2. 解析 Authorization 格式（Bearer Token）
	tokenStr, err := util.ParseBearerToken(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// 3. 验证 Token 并获取 UserID
	claims, err := util.ValidateToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的令牌", "detail": err.Error()})
		return
	}

	// 4. 从 Claims 中提取 UserID
	userID := claims.UserID

	// 5. 绑定请求体（使用提取的 UserID 替换请求体中的 UserID，防止篡改）
	var req dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数格式错误", "detail": err.Error()})
		return
	}
	req.UserID = userID // 强制使用 Token 中的 UserID，避免请求体伪造

	// 6. 后续业务逻辑（与原代码一致）
	orderId, err := h.idGen.NextID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "雪花 ID 生成失败", "detail": err.Error()})
		return
	}

	totalPrice := 0
	for _, orderItem := range req.CreateOrderItems {
		totalPrice += orderItem.ItemPrice * orderItem.ItemCount
	}

	if err := h.orderService.CreateOrder(orderId, req.UserID, totalPrice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建订单失败", "detail": err.Error()})
		return
	}

	if err := h.orderItemService.CreateOrderItem(orderId, req.UserID, req.CreateOrderItems); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建订单项失败", "detail": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"ok": 200, "order_id": orderId})
}
