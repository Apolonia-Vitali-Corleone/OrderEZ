package model

type CartDetail struct {
	CartDetailID int64 `gorm:"primaryKey;not null" json:"cart_detail_id"`
	CartID       int64 `gorm:"not null" json:"cart_id"`
	GoodID       int64 `gorm:"not null" json:"good_id"`
	Count        int   `gorm:"not null" json:"count"`
}

func (CartDetail) TableName() string {
	return "oe_cart_detail"
}
