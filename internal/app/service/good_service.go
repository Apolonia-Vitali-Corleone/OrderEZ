package service

import (
	"OrderEZ/internal/app/repository"
	"OrderEZ/internal/infrastructure/messaging"
	"OrderEZ/internal/po"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// GoodService 商品服务结构体
type GoodService struct {
	goodRepo *repository.GoodRepository
	redis    *redis.Client
	rabbitMQ *messaging.RabbitMQ
}

// NewGoodService 创建一个新的 GoodService 实例
func NewGoodService(db *gorm.DB, redisClient *redis.Client, mq *messaging.RabbitMQ) *GoodService {
	goodRepo := repository.NewGoodRepository(db)
	return &GoodService{
		goodRepo: goodRepo,
		redis:    redisClient,
		rabbitMQ: mq,
	}
}

// AddGood 方法用于添加一个新的商品
func (s *GoodService) AddGood(good po.Good) error {
	if err := s.goodRepo.Add(good); err != nil {
		return fmt.Errorf("failed to add good: %w", err)
	}

	// 清除商品列表缓存
	if err := s.clearGoodsCache(); err != nil {
		fmt.Printf("Failed to clear goods cache: %v\n", err)
	}

	// 发送消息到 RabbitMQ 进行异步处理，例如记录日志、更新统计信息等
	if err := s.rabbitMQ.Publish("good_added", []byte(fmt.Sprintf("Good added: %+v", good))); err != nil {
		fmt.Printf("Failed to publish message to RabbitMQ: %v\n", err)
	}

	return nil
}

// DeleteGood 根据商品 GoodID 删除商品
func (s *GoodService) DeleteGood(id int64) error {
	if err := s.goodRepo.DeleteGood(id); err != nil {
		return fmt.Errorf("failed to delete good: %w", err)
	}

	// 清除商品缓存
	if err := s.clearGoodCache(id); err != nil {
		fmt.Printf("Failed to clear good cache: %v\n", err)
	}

	// 清除商品列表缓存
	if err := s.clearGoodsCache(); err != nil {
		fmt.Printf("Failed to clear goods cache: %v\n", err)
	}

	// 发送消息到 RabbitMQ 进行异步处理
	if err := s.rabbitMQ.Publish("good_deleted", []byte(fmt.Sprintf("Good deleted: GoodID=%d", id))); err != nil {
		fmt.Printf("Failed to publish message to RabbitMQ: %v\n", err)
	}

	return nil
}

// UpdateGood 更新商品信息
func (s *GoodService) UpdateGood(good po.Good) error {
	if err := s.goodRepo.UpdateGood(good); err != nil {
		return fmt.Errorf("failed to update good: %w", err)
	}

	// 清除商品缓存
	if err := s.clearGoodCache(good.GoodID); err != nil {
		fmt.Printf("Failed to clear good cache: %v\n", err)
	}

	// 清除商品列表缓存
	if err := s.clearGoodsCache(); err != nil {
		fmt.Printf("Failed to clear goods cache: %v\n", err)
	}

	// 发送消息到 RabbitMQ 进行异步处理
	if err := s.rabbitMQ.Publish("good_updated", []byte(fmt.Sprintf("Good updated: %+v", good))); err != nil {
		fmt.Printf("Failed to publish message to RabbitMQ: %v\n", err)
	}

	return nil
}

// GetAllGoods 方法用于获取指定页码和每页数量的商品
func (s *GoodService) GetAllGoods(page, pageSize int) ([]po.Good, error) {
	// 尝试从 Redis 缓存中获取商品列表
	key := fmt.Sprintf("goods:page:%d:size:%d", page, pageSize)
	cachedGoods, err := s.getGoodsFromCache(key)
	if err == nil && len(cachedGoods) > 0 {
		return cachedGoods, nil
	}

	// 如果缓存中没有，从数据库中获取
	goods, err := s.goodRepo.GetAllGoods(page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to get goods from database: %w", err)
	}

	// 将从数据库获取的商品列表存入 Redis 缓存
	if err := s.setGoodsInCache(key, goods); err != nil {
		fmt.Printf("Failed to set goods in cache: %v\n", err)
	}

	return goods, nil
}

// GetGoodByID 根据商品 GoodID 获取商品信息
func (s *GoodService) GetGoodByID(id int64) (po.Good, error) {
	// 尝试从 Redis 缓存中获取商品信息
	key := fmt.Sprintf("good:id:%d", id)
	cachedGood, err := s.getGoodFromCache(key)
	if err == nil {
		return cachedGood, nil
	}

	// 如果缓存中没有，从数据库中获取
	good, err := s.goodRepo.GetGoodByID(id)
	if err != nil {
		return po.Good{}, fmt.Errorf("failed to get good by GoodID from database: %w", err)
	}

	// 将从数据库获取的商品信息存入 Redis 缓存
	if err := s.setGoodInCache(key, good); err != nil {
		fmt.Printf("Failed to set good in cache: %v\n", err)
	}

	return good, nil
}

/*
 * 这里是和redis相关的操作
 */

/*
getGoodsFromCache 从 Redis 缓存中获取商品列表
使用内容：
key := fmt.Sprintf("goods:page:%d:size:%d", page, pageSize)
cachedGoods, err := s.getGoodsFromCache(key)
*/
func (s *GoodService) getGoodsFromCache(key string) ([]po.Good, error) {
	val, err := s.redis.Get(s.redis.Context(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}
	var goods []po.Good
	err = json.Unmarshal([]byte(val), &goods)
	if err != nil {
		return nil, err
	}
	return goods, nil
}

// setGoodsInCache 将商品列表存入 Redis 缓存
// key就是goods:page:%d:size:%d，goods是查询出来的数据
func (s *GoodService) setGoodsInCache(key string, goods []po.Good) error {
	data, err := json.Marshal(goods)
	if err != nil {
		return err
	}
	return s.redis.Set(s.redis.Context(), key, string(data), 0).Err()
}

// getGoodFromCache 从 Redis 缓存中获取单个商品信息
func (s *GoodService) getGoodFromCache(key string) (po.Good, error) {
	val, err := s.redis.Get(s.redis.Context(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return po.Good{}, nil
		}
		return po.Good{}, err
	}
	var good po.Good
	err = json.Unmarshal([]byte(val), &good)
	if err != nil {
		return po.Good{}, err
	}
	return good, nil
}

// setGoodInCache 将单个商品信息存入 Redis 缓存
func (s *GoodService) setGoodInCache(key string, good po.Good) error {
	data, err := json.Marshal(good)
	if err != nil {
		return err
	}
	return s.redis.Set(s.redis.Context(), key, string(data), 0).Err()
}

// clearGoodCache 清除单个商品的缓存
func (s *GoodService) clearGoodCache(id int64) error {
	key := fmt.Sprintf("good:id:%d", id)
	return s.redis.Del(s.redis.Context(), key).Err()
}

// clearGoodsCache 清除商品列表的缓存
func (s *GoodService) clearGoodsCache() error {
	keys, err := s.redis.Keys(s.redis.Context(), "goods:*").Result()
	if err != nil {
		return err
	}
	for _, key := range keys {
		if err := s.redis.Del(s.redis.Context(), key).Err(); err != nil {
			return err
		}
	}
	return nil
}
