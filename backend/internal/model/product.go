package model

import "time"

// Product represents a product/drug in the inventory
type Product struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Stock     int       `json:"stock" db:"stock"`
	Price     float64   `json:"price" db:"price"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// CreateProductRequest represents the request body for creating a product
type CreateProductRequest struct {
	Name  string  `json:"name" binding:"required"`
	Stock int     `json:"stock" binding:"required,min=0"`
	Price float64 `json:"price" binding:"required,min=0"`
}

