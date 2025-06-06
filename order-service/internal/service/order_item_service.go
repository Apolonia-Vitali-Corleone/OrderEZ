package service

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"order-service/infrastructure/messaging"
	"order-service/internal/model/dto"
	"order-service/internal/model/po"
	"order-service/internal/repository"
	"order-service/util"
)

type OrderItemService struct {
	orderItemRepository *repository.OrderItemRepository
	redis               *redis.Client
	rabbitMQ            *messaging.RabbitMQ
	idGen               *util.Snowflake // ✅ 雪花 ID 生成器
}

func (s OrderItemService) CreateOrderItem(orderID int64, userID int64, createOrderItem []dto.CreateOrderItem) error {
	for _, createOrderItem := range createOrderItem {
		orderItemID, err := s.idGen.NextID()
		if err != nil {
			return fmt.Errorf("生成订单项目ID失败: %w", err)
		}

		orderItem := &po.OrderItem{
			OrderItemID: orderItemID,
			OrderID:     orderID,
			UserID:      userID,
			ItemID:      createOrderItem.ItemID,
			ItemName:    createOrderItem.ItemName,
			ItemPrice:   createOrderItem.ItemPrice,
			ItemCount:   createOrderItem.ItemCount,
		}
		err = s.orderItemRepository.CreateOrderItem(orderItem)
		if err != nil {
			return fmt.Errorf("创建订单项目失败: %w", err)
		}
	}
	return nil
}

func NewOrderItemService(db *gorm.DB, redisClient *redis.Client, mq *messaging.RabbitMQ) *OrderItemService {
	orderItemRepository := repository.NewOrderItemRepository(db)
	idGen, err := util.NewSnowflake(1, 1)
	if err != nil {
		panic("failed to init snowflake: " + err.Error())
	}
	return &OrderItemService{
		orderItemRepository: orderItemRepository,
		redis:               redisClient,
		rabbitMQ:            mq,
		idGen:               idGen,
	}
}

//func (s *OrderService) CreateOrder(req *dto.CreateOrderRequest) (*vo.OrderVO, error) {
//	orderID, err := s.idGen.NextID()
//
//	order := &model.Order{OrderID: orderID, UserID: req.UserID}
//
//	err = s.orderRepo.CreateOrder(order)
//
//	if err != nil {
//		return nil, err
//	}
//	return nil, nil
//}

//func (s *OrderService) CreateOrder(req *model.CreateOrderRequest) (*model.OrderVO, error)
//{

//
//	var totalPrice int
//	var orderItems []model.OrderDetail
//
//	// 校验商品，累加总价，构建订单详情
//	for _, item := range req.Items {
//		good, err := s.orderRepo.GetGoodByID(item.GoodID)
//		if err != nil {
//			return nil, fmt.Errorf("商品 %d 查询失败: %w", item.GoodID, err)
//		}
//		if good.GoodStock < item.Amount {
//			return nil, fmt.Errorf("商品 %s 库存不足", good.GoodName)
//		}
//
//		// 累加价格
//		totalPrice += good.GoodPrice * item.Amount
//
//		// 构造订单详情项
//		orderItems = append(orderItems, model.OrderDetail{
//			OrderID: orderID,
//			GoodID:  good.GoodID,
//			Amount:  item.Amount,
//			Price:   good.GoodPrice,
//		})
//	}
//
//	order := &model.Order{
//		OrderID:     orderID,
//		UserID:      req.UserID,
//		OrderNumber: fmt.Sprintf("OE-%d", orderID),
//		TotalPrice:  totalPrice,
//	}
//
//	// 事务保存订单和订单详情
//	if err := s.orderRepo.CreateOrderWithDetails(order, orderItems); err != nil {
//		return nil, err
//	}
//
//	// 返回 VO
//	orderVO := &vo.OrderVO{
//		OrderID:     order.OrderID,
//		OrderNumber: order.OrderNumber,
//		TotalPrice:  fmt.Sprintf("¥%.2f", float64(order.TotalPrice)/100),
//		StatusLabel: "已下单",
//	}
//
//	return orderVO, nil
//}
//
//// 其他订单服务方法...
