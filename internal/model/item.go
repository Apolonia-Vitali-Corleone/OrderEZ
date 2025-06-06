package model

type Item struct {
	ItemID    int64  `gorm:"primaryKey;not null" json:"item_id"`
	ItemName  string `gorm:"not null" json:"item_name"`
	ItemPrice int    `gorm:"not null" json:"item_price"`
	ItemStock int    `gorm:"not null" json:"item_stock"`
}

func (Item) TableName() string {
	return "oe_item"
}
