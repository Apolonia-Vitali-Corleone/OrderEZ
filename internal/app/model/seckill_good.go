package model

type SeckillGood struct {
	SeckillGoodID    string  `gorm:"primaryKey;not null" json:"seckill_good_id"`
	SeckillGoodName  string  `gorm:"not null" json:"seckill_good_name"`
	SeckillGoodStock int     `gorm:"not null" json:"seckill_good_stock"`
	SeckillGoodPrice float64 `gorm:"not null" json:"seckill_good_price"`
}

func (SeckillGood) TableName() string {
	return "oe_seckill_good"
}
