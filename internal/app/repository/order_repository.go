package repository

import (
	"OrderEZ/internal/app/model"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(order *model.Order) error {
	return r.db.Create(order).Error
}

// 其他订单操作方法...
