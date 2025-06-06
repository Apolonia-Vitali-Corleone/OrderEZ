package repository

import (
	"fmt"
	"gorm.io/gorm"
	"order-service/internal/model/po"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(order *po.Order) error {
	result := r.db.Create(order)

	// 检查是否有错误发生
	if result.Error != nil {
		// 构建更详细的错误信息
		return fmt.Errorf("创建订单失败: %w", result.Error)
	}

	// 检查是否成功创建了记录
	if result.RowsAffected == 0 {
		return fmt.Errorf("创建订单失败: 没有记录被插入")
	}

	return nil
}
