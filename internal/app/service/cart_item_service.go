package service

import (
	"OrderEZ/internal/app/repository"
	"OrderEZ/internal/infrastructure/messaging"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// CartItemService 定义购物车详情服务结构体
type CartItemService struct {
	cartDetailRepo *repository.CartItemRepository
	redis          *redis.Client
	rabbitMQ       *messaging.RabbitMQ
}

// NewCartDetailService 创建一个新的 CartItemService 实例
func NewCartDetailService(db *gorm.DB, redisClient *redis.Client, mq *messaging.RabbitMQ) *CartItemService {
	cartDetailRepository := repository.NewCartItemRepository(db)
	return &CartItemService{
		cartDetailRepo: cartDetailRepository,
		redis:          redisClient,
		rabbitMQ:       mq,
	}
}

//// AddCartItem adds a new item to the cart
//func (s *CartItemService) AddCartItem(cartID, itemID int64) (*model.CartItem, error) {
//	cartItem := &model.CartItem{
//		CartID: cartID,
//		ItemID: itemID,
//	}
//	if err := s.repo.CreateCartItem(cartItem); err != nil {
//		return nil, err
//	}
//	return cartItem, nil
//}
//
//// GetCartItem retrieves a cart item by its ID
//func (s *CartItemService) GetCartItem(cartItemID int64) (*model.CartItem, error) {
//	cartItem, err := s.repo.GetCartItemByID(cartItemID)
//	if err != nil {
//		return nil, err
//	}
//	return cartItem, nil
//}
//
//// GetCartItemsByCartID retrieves all items in a cart
//func (s *CartItemService) GetCartItemsByCartID(cartID int64) ([]model.CartItem, error) {
//	cartItems, err := s.repo.GetCartItemsByCartID(cartID)
//	if err != nil {
//		return nil, err
//	}
//	return cartItems, nil
//}
//
//// UpdateCartItem updates a cart item
//func (s *CartItemService) UpdateCartItem(cartItemID, newCartID, newItemID int64) (*model.CartItem, error) {
//	cartItem, err := s.repo.GetCartItemByID(cartItemID)
//	if err != nil {
//		return nil, err
//	}
//	cartItem.CartID = newCartID
//	cartItem.ItemID = newItemID
//	if err := s.repo.UpdateCartItem(cartItem); err != nil {
//		return nil, err
//	}
//	return cartItem, nil
//}
//
//// DeleteCartItem removes a cart item
//func (s *CartItemService) DeleteCartItem(cartItemID int64) error {
//	return s.repo.DeleteCartItem(cartItemID)
//}
//
//// DeleteCartItemsByCartID removes all items from a cart
//func (s *CartItemService) DeleteCartItemsByCartID(cartID int64) error {
//	return s.repo.DeleteCartItemsByCartID(cartID)
//}
