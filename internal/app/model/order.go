package model

type Order struct {
	OrderID     uint   `gorm:"primaryKey;not null" json:"order_id"`
	UserID      uint   `gorm:"unique;not null" json:"user_id"`
	OrderNumber string `gorm:"unique;not null" json:"order_number"`
	TotalPrice  int    `gorm:"not null" json:"total_price"`
}

func (Order) TableName() string {
	return "oe_order"
}
