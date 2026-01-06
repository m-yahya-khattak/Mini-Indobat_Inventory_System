package service

import (
	"context"
	"mini-indobat-inventory/internal/model"
	"mini-indobat-inventory/internal/repository"
)

// ProductService handles product business logic
type ProductService struct {
	productRepo *repository.ProductRepository
}

// NewProductService creates a new product service
func NewProductService(productRepo *repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

// GetAllProducts retrieves all products
func (s *ProductService) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	return s.productRepo.GetAll(ctx)
}

// CreateProduct creates a new product
func (s *ProductService) CreateProduct(ctx context.Context, req model.CreateProductRequest) (*model.Product, error) {
	// Business logic validation
	if req.Name == "" {
		return nil, ErrInvalidInput
	}
	if req.Stock < 0 {
		return nil, ErrInvalidStock
	}
	if req.Price < 0 {
		return nil, ErrInvalidPrice
	}

	return s.productRepo.Create(ctx, req)
}

// GetProductByID retrieves a product by ID
func (s *ProductService) GetProductByID(ctx context.Context, id int) (*model.Product, error) {
	return s.productRepo.GetByID(ctx, id)
}

