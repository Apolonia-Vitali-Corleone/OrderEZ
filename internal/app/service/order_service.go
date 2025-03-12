package service

import (
	"OrderEZ/internal/app/model"
	"OrderEZ/internal/app/repository"
	"OrderEZ/internal/infrastructure/messaging"
	"encoding/json"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

type OrderService struct {
	orderRepo *repository.OrderRepository
	rabbitMQ  *messaging.RabbitMQ
}

func NewOrderService(db *gorm.DB, rabbitMQConn *amqp.Connection) *OrderService {
	orderRepo := repository.NewOrderRepository(db)
	rabbitMQ := messaging.NewRabbitMQ(rabbitMQConn)
	return &OrderService{
		orderRepo: orderRepo,
		rabbitMQ:  rabbitMQ,
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
