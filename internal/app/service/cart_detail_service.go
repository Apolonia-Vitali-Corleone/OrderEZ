package service

import (
	"OrderEZ/internal/app/model"
	"OrderEZ/internal/app/repository"
	"OrderEZ/internal/infrastructure/messaging"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// CartDetailService 定义购物车详情服务结构体
type CartDetailService struct {
	cartDetailRepo *repository.CartDetailRepository
	redis          *redis.Client
	rabbitMQ       *messaging.RabbitMQ
}

// NewCartDetailService 创建一个新的 CartDetailService 实例
func NewCartDetailService(db *gorm.DB, redisClient *redis.Client, mq *messaging.RabbitMQ) *CartDetailService {
	cartDetailRepository := repository.NewCartDetailRepository(db)
	return &CartDetailService{
		cartDetailRepo: cartDetailRepository,
		redis:          redisClient,
		rabbitMQ:       mq,
	}
}

// GetCartDetailListByCartID 根据 cart_id 获取购物车详情列表
func (s *CartDetailService) GetCartDetailListByCartID(cartID int64) ([]model.CartDetail, error) {
	return s.cartDetailRepo.GetCartDetailListByCartID(cartID)
}
