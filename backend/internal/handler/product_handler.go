package handler

import (
	"errors"
	"net/http"
	"mini-indobat-inventory/internal/model"
	"mini-indobat-inventory/internal/service"

	"github.com/gin-gonic/gin"
)

// ProductHandler handles product HTTP requests
type ProductHandler struct {
	productService *service.ProductService
}

// NewProductHandler creates a new product handler
func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// GetProducts handles GET /products
func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.productService.GetAllProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// CreateProduct handles POST /products
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req model.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	product, err := h.productService.CreateProduct(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidInput) || errors.Is(err, service.ErrInvalidStock) || errors.Is(err, service.ErrInvalidPrice) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

