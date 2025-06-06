package model

type Cart struct {
	CartID int64 `gorm:"primaryKey;not null" json:"cart_id"`
	UserID int64 `gorm:"not null;unique" json:"user_id"`
}

func (Cart) TableName() string {
	return "oe_cart"
}
