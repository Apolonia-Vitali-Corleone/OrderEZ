package po

type Order struct {
	OrderID     int64  `gorm:"column:order_id;primaryKey;not null;autoIncrement:false" json:"order_id"`
	UserID      int64  `gorm:"column:user_id;unique;not null" json:"user_id"`
	OrderNumber string `gorm:"column:order_number;unique;not null" json:"order_number"`
	TotalPrice  int    `gorm:"column:total_price;not null" json:"total_price"`
}

func (Order) TableName() string {
	return "oe_order"
}
