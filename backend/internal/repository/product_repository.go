package repository

import (
	"context"
	"mini-indobat-inventory/internal/config"
	"mini-indobat-inventory/internal/model"

	"github.com/jackc/pgx/v5"
)

// ProductRepository handles product data access
type ProductRepository struct{}

// NewProductRepository creates a new product repository
func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

// GetAll retrieves all products
func (r *ProductRepository) GetAll(ctx context.Context) ([]model.Product, error) {
	query := `
		SELECT id, name, stock, price, created_at, updated_at
		FROM products
		ORDER BY id
	`

	rows, err := config.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Stock,
			&p.Price,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// GetByID retrieves a product by ID
func (r *ProductRepository) GetByID(ctx context.Context, id int) (*model.Product, error) {
	query := `
		SELECT id, name, stock, price, created_at, updated_at
		FROM products
		WHERE id = $1
	`

	var p model.Product
	err := config.DB.QueryRow(ctx, query, id).Scan(
		&p.ID,
		&p.Name,
		&p.Stock,
		&p.Price,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

// GetByIDWithLock retrieves a product by ID with row-level lock (for transactions)
// This uses SELECT FOR UPDATE to prevent race conditions
func (r *ProductRepository) GetByIDWithLock(ctx context.Context, tx pgx.Tx, id int) (*model.Product, error) {
	query := `
		SELECT id, name, stock, price, created_at, updated_at
		FROM products
		WHERE id = $1
		FOR UPDATE
	`

	var p model.Product
	var err error

	// Use transaction if provided, otherwise use connection pool
	if tx != nil {
		err = tx.QueryRow(ctx, query, id).Scan(
			&p.ID,
			&p.Name,
			&p.Stock,
			&p.Price,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
	} else {
		err = config.DB.QueryRow(ctx, query, id).Scan(
			&p.ID,
			&p.Name,
			&p.Stock,
			&p.Price,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
	}

	if err != nil {
		return nil, err
	}

	return &p, nil
}

// Create creates a new product
func (r *ProductRepository) Create(ctx context.Context, req model.CreateProductRequest) (*model.Product, error) {
	query := `
		INSERT INTO products (name, stock, price)
		VALUES ($1, $2, $3)
		RETURNING id, name, stock, price, created_at, updated_at
	`

	var p model.Product
	err := config.DB.QueryRow(ctx, query, req.Name, req.Stock, req.Price).Scan(
		&p.ID,
		&p.Name,
		&p.Stock,
		&p.Price,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

// UpdateStock updates the stock of a product (used in transactions)
func (r *ProductRepository) UpdateStock(ctx context.Context, tx pgx.Tx, productID int, newStock int) error {
	query := `
		UPDATE products
		SET stock = $1
		WHERE id = $2
	`

	var err error
	if tx != nil {
		_, err = tx.Exec(ctx, query, newStock, productID)
	} else {
		_, err = config.DB.Exec(ctx, query, newStock, productID)
	}

	return err
}

