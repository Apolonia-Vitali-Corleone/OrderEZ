package repository

import (
	"OrderEZ/internal/po"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(order *po.Order) error {
	return r.db.Create(order).Error
}

// 其他订单操作方法...
