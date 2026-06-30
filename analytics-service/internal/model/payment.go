package model

type Payment struct {
	OrderID int64  `json:"order_id"`
	Status  string `json:"status"`
}
