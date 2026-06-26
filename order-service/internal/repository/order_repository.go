package repository

import (
	"database/sql"

	"github.com/madhurima877/order-service/internal/model"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Create(order *model.Order) error {
	query := `INSERT INTO orders (user_id, product, quantity, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id `
	return r.db.QueryRow(query, order.UserID,
		order.Product,
		order.Quantity,
		order.Status).Scan(&order.ID)
}
