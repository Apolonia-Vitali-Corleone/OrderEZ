package repository

import (
	"OrderEZ/internal/app/model"
	"gorm.io/gorm"
)

// GoodRepository 定义商品仓库结构体
type GoodRepository struct {
	db *gorm.DB
}

// NewGoodRepository 创建一个新的 GoodRepository 实例
func NewGoodRepository(db *gorm.DB) *GoodRepository {
	return &GoodRepository{db: db}
}

// GetAllGoods 查询指定页码和每页数量的商品
func (r *GoodRepository) GetAllGoods(page, pageSize int) ([]model.Good, error) {
	var goods []model.Good
	offset := (page - 1) * pageSize
	result := r.db.Offset(offset).Limit(pageSize).Find(&goods)
	if result.Error != nil {
		return nil, result.Error
	}
	return goods, nil
}

// Add 方法用于向数据库中添加一个新的商品
func (r *GoodRepository) Add(good model.Good) error {
	result := r.db.Create(&good)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteGood 根据商品 GoodID 删除商品
func (r *GoodRepository) DeleteGood(id int64) error {
	result := r.db.Delete(&model.Good{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateGood 更新商品信息
func (r *GoodRepository) UpdateGood(good model.Good) error {
	result := r.db.Save(&good)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetGoodByID 根据商品 GoodID 获取商品信息
func (r *GoodRepository) GetGoodByID(id int64) (model.Good, error) {
	var good model.Good
	result := r.db.First(&good, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return model.Good{}, nil
		}
		return model.Good{}, result.Error
	}
	return good, nil
}
