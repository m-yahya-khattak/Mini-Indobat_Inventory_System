package service

import (
	"context"
	"errors"
	"fmt"
	"mini-indobat-inventory/internal/config"
	"mini-indobat-inventory/internal/model"
	"mini-indobat-inventory/internal/repository"

	"github.com/jackc/pgx/v5"
)

var (
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrProductNotFound   = errors.New("product not found")
	ErrInvalidQuantity   = errors.New("invalid quantity")
	ErrInvalidDiscount   = errors.New("invalid discount percentage")
)

// OrderService handles order business logic
type OrderService struct {
	orderRepo   *repository.OrderRepository
	productRepo *repository.ProductRepository
}

// NewOrderService creates a new order service
func NewOrderService(orderRepo *repository.OrderRepository, productRepo *repository.ProductRepository) *OrderService {
	return &OrderService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

// CreateOrder creates a new order with race condition protection using database transactions.
func (s *OrderService) CreateOrder(ctx context.Context, req model.CreateOrderRequest) (*model.OrderResponse, error) {
	if req.Quantity <= 0 {
		return nil, ErrInvalidQuantity
	}
	if req.DiscountPercent < 0 || req.DiscountPercent > 100 {
		return nil, ErrInvalidDiscount
	}

	tx, err := config.DB.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	product, err := s.productRepo.GetByIDWithLock(ctx, tx, req.ProductID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrProductNotFound
		}
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	if product.Stock < req.Quantity {
		return nil, fmt.Errorf("%w: available stock is %d, requested %d", ErrInsufficientStock, product.Stock, req.Quantity)
	}

	totalPrice := CalculateOrderPrice(product.Price, req.Quantity, req.DiscountPercent)

	order := &model.Order{
		ProductID:      req.ProductID,
		Quantity:       req.Quantity,
		DiscountPercent: req.DiscountPercent,
		TotalPrice:     totalPrice,
	}

	if err := s.orderRepo.Create(ctx, tx, order); err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	newStock := product.Stock - req.Quantity
	if err := s.productRepo.UpdateStock(ctx, tx, req.ProductID, newStock); err != nil {
		return nil, fmt.Errorf("failed to update stock: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &model.OrderResponse{
		OrderID:    order.ID,
		ProductID:  order.ProductID,
		Quantity:   order.Quantity,
		TotalPrice: order.TotalPrice,
		Message:    "Order successful",
	}, nil
}

