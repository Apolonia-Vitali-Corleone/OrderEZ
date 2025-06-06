package repository

import (
	"fmt"
	"gorm.io/gorm"
	"order-service/internal/model/po"
)

type OrderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *OrderItemRepository {
	return &OrderItemRepository{db: db}
}

func (r OrderItemRepository) CreateOrderItem(orderItem *po.OrderItem) error {
	result := r.db.Create(orderItem)
	if result.Error != nil {
		return fmt.Errorf("创建订单项目失败: %w", result.Error)
	}
	return nil
}
