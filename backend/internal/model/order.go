package model

import "time"

// Order represents a purchase transaction
type Order struct {
	ID             int       `json:"id" db:"id"`
	ProductID      int       `json:"product_id" db:"product_id"`
	Quantity       int       `json:"quantity" db:"quantity"`
	DiscountPercent float64  `json:"discount_percent" db:"discount_percent"`
	TotalPrice     float64   `json:"total_price" db:"total_price"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// CreateOrderRequest represents the request body for creating an order
type CreateOrderRequest struct {
	ProductID      int     `json:"product_id" binding:"required"`
	Quantity       int     `json:"quantity" binding:"required,min=1"`
	DiscountPercent float64 `json:"discount_percent" binding:"required,min=0,max=100"`
}

// OrderResponse represents the response after creating an order
type OrderResponse struct {
	OrderID    int     `json:"order_id"`
	ProductID  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	Message    string  `json:"message"`
}

