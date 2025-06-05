package repository

import (
	"OrderEZ/internal/po"
	"gorm.io/gorm"
)

// CartRepository 定义购物车结构体
type CartRepository struct {
	db *gorm.DB
}

// NewCartRepository 创建一个新的 CartRepository 实例
func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

// GetCartIDByUserID 根据 user_id 获取 cart_id
func (r *CartRepository) GetCartIDByUserID(userID int64) (int64, error) {
	var cart po.Cart
	err := r.db.Select("cart_id").Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		return 0, err
	}
	return cart.CartID, nil
}
