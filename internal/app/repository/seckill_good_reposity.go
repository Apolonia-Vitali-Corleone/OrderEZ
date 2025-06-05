package repository

import (
	"OrderEZ/internal/po"
	"gorm.io/gorm"
)

// SeckillGoodRepository 定义商品仓库结构体
type SeckillGoodRepository struct {
	db *gorm.DB
}

// NewSeckillGoodRepository 创建一个新的 GoodRepository 实例
func NewSeckillGoodRepository(db *gorm.DB) *SeckillGoodRepository {
	return &SeckillGoodRepository{db: db}
}

// GetAllSeckillGoods 查询指定页码和每页数量的商品
func (r *SeckillGoodRepository) GetAllSeckillGoods(page, pageSize int) ([]po.SeckillGood, error) {
	var seckillGood []po.SeckillGood
	offset := (page - 1) * pageSize
	result := r.db.Offset(offset).Limit(pageSize).Find(&seckillGood)
	if result.Error != nil {
		return nil, result.Error
	}
	return seckillGood, nil
}

//// Add 方法用于向数据库中添加一个新的商品
//func (r *GoodRepository) Add(good po.Good) error {
//	result := r.db.Create(&good)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//// DeleteGood 根据商品 GoodID 删除商品
//func (r *GoodRepository) DeleteGood(id int64) error {
//	result := r.db.Delete(&po.Good{}, id)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//// UpdateGood 更新商品信息
//func (r *GoodRepository) UpdateGood(good po.Good) error {
//	result := r.db.Save(&good)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//// GetGoodByID 根据商品 GoodID 获取商品信息
//func (r *GoodRepository) GetGoodByID(id int64) (po.Good, error) {
//	var good po.Good
//	result := r.db.First(&good, id)
//	if result.Error != nil {
//		if result.Error == gorm.ErrRecordNotFound {
//			return po.Good{}, nil
//		}
//		return po.Good{}, result.Error
//	}
//	return good, nil
//}
