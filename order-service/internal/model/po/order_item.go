package po

type OrderItem struct {
	OrderItemID int64  `gorm:"primaryKey;not null" json:"order_item_id"`
	OrderID     int64  `gorm:"not null" json:"order_id"`
	UserID      int64  `gorm:"not null" json:"user_id"`
	ItemID      int64  `gorm:"not null" json:"item_id"`
	ItemName    string `gorm:"not null" json:"item_name"`
	ItemPrice   int    `gorm:"not null" json:"item_price"`
	ItemCount   int    `gorm:"not null" json:"item_count"`
}

func (OrderItem) TableName() string {
	return "oe_order_item"
}
