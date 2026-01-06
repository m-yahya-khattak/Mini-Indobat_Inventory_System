package repository

import (
	"context"
	"mini-indobat-inventory/internal/config"
	"mini-indobat-inventory/internal/model"

	"github.com/jackc/pgx/v5"
)

// OrderRepository handles order data access
type OrderRepository struct{}

// NewOrderRepository creates a new order repository
func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

// Create creates a new order
func (r *OrderRepository) Create(ctx context.Context, tx pgx.Tx, order *model.Order) error {
	query := `
		INSERT INTO orders (product_id, quantity, discount_percent, total_price)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`

	var err error
	if tx != nil {
		err = tx.QueryRow(ctx, query,
			order.ProductID,
			order.Quantity,
			order.DiscountPercent,
			order.TotalPrice,
		).Scan(&order.ID, &order.CreatedAt)
	} else {
		err = config.DB.QueryRow(ctx, query,
			order.ProductID,
			order.Quantity,
			order.DiscountPercent,
			order.TotalPrice,
		).Scan(&order.ID, &order.CreatedAt)
	}

	return err
}

