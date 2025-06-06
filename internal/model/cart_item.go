package model

type CartItem struct {
	CartItemID int64 `gorm:"primaryKey;not null" json:"cart_item_id"`
	CartID     int64 `gorm:"not null" json:"cart_id"`
	ItemID     int64 `gorm:"not null" json:"item_id"`
}

func (CartItem) TableName() string {
	return "oe_cart_item"
}
