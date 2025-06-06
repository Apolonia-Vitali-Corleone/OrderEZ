package repository

import (
	"OrderEZ/internal/model"
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
	var cart model.Cart
	err := r.db.Select("cart_id").Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		return 0, err
	}
	return cart.CartID, nil
}

// CreateCart 用户注册成功，创建对应的购物车
func (r *CartRepository) CreateCart(cart *model.Cart) error {
	return r.db.Create(cart).Error
}

// DeleteCart 用户注销后，删除对应的购物车
func (r *CartRepository) DeleteCart(userID int64) error {
	return r.db.Where("user_id = ?", userID).Delete(&model.Cart{}).Error
}
