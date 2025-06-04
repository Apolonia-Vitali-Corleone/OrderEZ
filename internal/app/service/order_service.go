package service

import (
	"OrderEZ/internal/app/model"
	"OrderEZ/internal/app/repository"
	"OrderEZ/internal/infrastructure/messaging"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type OrderService struct {
	orderRepo *repository.OrderRepository
	redis     *redis.Client
	rabbitMQ  *messaging.RabbitMQ
}

func NewOrderService(db *gorm.DB, redisClient *redis.Client, mq *messaging.RabbitMQ) *OrderService {
	orderRepo := repository.NewOrderRepository(db)
	return &OrderService{
		orderRepo: orderRepo,
		redis:     redisClient,
		rabbitMQ:  mq,
	}
}

func (s *OrderService) CreateOrder(order *model.Order) error {
	// 保存订单到数据库
	if err := s.orderRepo.CreateOrder(order); err != nil {
		return err
	}

	// 发送订单消息到 RabbitMQ 进行异步处理
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return err
	}
	return s.rabbitMQ.Publish("order_queue", orderJSON)
}

// 其他订单服务方法...
