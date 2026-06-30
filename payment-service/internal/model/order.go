package model

type Order struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
	Status   string `json:"status"`
}