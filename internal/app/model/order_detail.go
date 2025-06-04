package model

type OrderDetail struct {
	OrderDetailID int    `gorm:"primaryKey;not null" json:"order_detail_id"`
	OrderID       int    `gorm:"not null" json:"order_id"`
	UserID        int    `gorm:"not null" json:"user_id"`
	GoodID        int    `gorm:"not null" json:"good_id"`
	GoodName      string `gorm:"not null" json:"good_name"`
	GoodPrice     int    `gorm:"not null" json:"good_price"`
	GoodCount     int    `gorm:"not null" json:"good_count"`
}

func (OrderDetail) TableName() string {
	return "oe_order_detail"
}
