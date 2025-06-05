package dto

type CreateOrderRequest struct {
	UserID      int64  `json:"user_id"`
	OrderNumber string `json:"order_number"`
	TotalPrice  int    `json:"total_price"`
}
