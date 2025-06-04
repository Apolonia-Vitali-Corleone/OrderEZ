package model

type Good struct {
	GoodID    int64  `gorm:"primaryKey;not null" json:"good_id"`
	GoodName  string `gorm:"not null" json:"good_name"`
	GoodPrice int    `gorm:"not null" json:"good_price"`
	GoodStock int    `gorm:"not null" json:"good_stock"`
}

func (Good) TableName() string {
	return "oe_good"
}
