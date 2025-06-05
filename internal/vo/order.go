package vo

type OrderVO struct {
	OrderID     int64  `json:"order_id"`
	OrderNumber string `json:"order_number"`
	TotalPrice  string `json:"total_price"`
	StatusLabel string `json:"status"`
}
