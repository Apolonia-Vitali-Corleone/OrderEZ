package model

type SeckillItem struct {
	SeckillItemID    string `gorm:"primaryKey;not null" json:"seckill_item_id"`
	SeckillItemName  string `gorm:"not null" json:"seckill_item_name"`
	SeckillItemPrice int    `gorm:"not null" json:"seckill_item_price"`
	SeckillItemStock int    `gorm:"not null" json:"seckill_item_stock"`
}

func (SeckillItem) TableName() string {
	return "oe_seckill_item"
}
