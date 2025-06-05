package service

import (
	"OrderEZ/internal/app/repository"
	"OrderEZ/internal/dto"
	"OrderEZ/internal/infrastructure/messaging"
	"OrderEZ/internal/po"
	"OrderEZ/internal/util"
	"OrderEZ/internal/vo"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type OrderService struct {
	orderRepo *repository.OrderRepository
	redis     *redis.Client
	rabbitMQ  *messaging.RabbitMQ
	idGen     *util.Snowflake // ✅ 雪花 ID 生成器
}

func NewOrderService(db *gorm.DB, redisClient *redis.Client, mq *messaging.RabbitMQ) *OrderService {
	orderRepo := repository.NewOrderRepository(db)
	idGen, err := util.NewSnowflake(1, 1)
	if err != nil {
		panic("failed to init snowflake: " + err.Error())
	}
	return &OrderService{
		orderRepo: orderRepo,
		redis:     redisClient,
		rabbitMQ:  mq,
		idGen:     idGen,
	}
}

func (s *OrderService) CreateOrder(req *dto.CreateOrderRequest) (*vo.OrderVO, error) {
	id, err := s.idGen.NextID()
	if err != nil {
		return nil, err
	}

	order := &po.Order{
		OrderID:     id,
		UserID:      req.UserID,
		OrderNumber: req.OrderNumber,
		TotalPrice:  req.TotalPrice,
	}

	// 保存订单
	if err := s.orderRepo.CreateOrder(order); err != nil {
		return nil, err
	}

	//// 发送订单消息到 MQ
	//orderJSON, err := json.Marshal(order)
	//if err != nil {
	//	return nil, err
	//}
	//if err := s.rabbitMQ.Publish("order_queue", orderJSON); err != nil {
	//	return nil, err
	//}

	// 构造返回 VO
	orderVO := &vo.OrderVO{
		OrderID:     order.OrderID,
		OrderNumber: order.OrderNumber,
		TotalPrice:  fmt.Sprintf("¥%.2f", float64(order.TotalPrice)/100),
		StatusLabel: "已下单", // 你可以根据状态扩展
	}
	return orderVO, nil
}

// 其他订单服务方法...
