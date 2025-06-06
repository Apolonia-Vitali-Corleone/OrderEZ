package service

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"order-service/infrastructure/messaging"
	"order-service/internal/model/po"
	"order-service/internal/repository"
	"order-service/util"
)

type OrderService struct {
	orderRepository *repository.OrderRepository
	redis           *redis.Client
	rabbitMQ        *messaging.RabbitMQ
	idGen           *util.Snowflake // ✅ 雪花 ID 生成器
}

func (s OrderService) CreateOrder(id int64, userID int64, totalPrice int) error {
	err := s.orderRepository.CreateOrder(&po.Order{OrderID: id, UserID: userID, TotalPrice: totalPrice})
	if err != nil {
		return err
	}
	return nil
}

func NewOrderService(db *gorm.DB, redisClient *redis.Client, mq *messaging.RabbitMQ) *OrderService {
	orderRepository := repository.NewOrderRepository(db)
	idGen, err := util.NewSnowflake(1, 1)
	if err != nil {
		panic("failed to init snowflake: " + err.Error())
	}
	return &OrderService{
		orderRepository: orderRepository,
		redis:           redisClient,
		rabbitMQ:        mq,
		idGen:           idGen,
	}
}
