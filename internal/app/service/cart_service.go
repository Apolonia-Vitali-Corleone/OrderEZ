package service

import (
	"OrderEZ/internal/app/repository"
	"OrderEZ/internal/infrastructure/messaging"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// CartService 商品服务结构体
type CartService struct {
	cartRepo *repository.CartRepository
	redis    *redis.Client
	rabbitMQ *messaging.RabbitMQ
}

// NewCartService 创建一个新的 CartService 实例
func NewCartService(db *gorm.DB, redisClient *redis.Client, mq *messaging.RabbitMQ) *CartService {
	cartRepo := repository.NewCartRepository(db)
	return &CartService{
		cartRepo: cartRepo,
		redis:    redisClient,
		rabbitMQ: mq,
	}
}

// GetCartIDByUserID 根据 user_id 获取 cart_id
func (s *CartService) GetCartIDByUserID(userID int64) (int64, error) {
	return s.cartRepo.GetCartIDByUserID(userID)
}
