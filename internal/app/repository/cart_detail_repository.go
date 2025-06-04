package repository

import (
	"OrderEZ/internal/app/model" // 请替换为实际的 model 包路径
	"gorm.io/gorm"
)

// CartDetailRepository 定义购物车结构体
type CartDetailRepository struct {
	db *gorm.DB
}

// NewCartDetailRepository 创建一个新的 CartDetailRepository 实例
func NewCartDetailRepository(db *gorm.DB) *CartDetailRepository {
	return &CartDetailRepository{db: db}
}

// GetCartDetailListByCartID 根据 cart_id 获取购物车详情列表
func (r *CartDetailRepository) GetCartDetailListByCartID(cartID int64) ([]model.CartDetail, error) {
	var cartDetails []model.CartDetail
	result := r.db.Where("cart_id = ?", cartID).Find(&cartDetails)
	if result.Error != nil {
		return nil, result.Error
	}
	return cartDetails, nil
}
