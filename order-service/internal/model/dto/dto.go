package dto

// CreateOrderRequest 表示前端创建订单时传来的请求体
type CreateOrderRequest struct {
	UserID           int64             `json:"user_id"`
	CreateOrderItems []CreateOrderItem `json:"create_order_items"`
}

type CreateOrderItem struct {
	ItemID    int64  `json:"item_id"`
	ItemName  string `json:"item_name"`
	ItemPrice int    `json:"item_price"`
	ItemCount int    `json:"item_count"`
}
