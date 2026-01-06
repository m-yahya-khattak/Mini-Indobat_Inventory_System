package handler

import (
	"errors"
	"net/http"
	"mini-indobat-inventory/internal/model"
	"mini-indobat-inventory/internal/service"

	"github.com/gin-gonic/gin"
)

// OrderHandler handles order HTTP requests
type OrderHandler struct {
	orderService *service.OrderService
}

// NewOrderHandler creates a new order handler
func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

// CreateOrder handles POST /order
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req model.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	response, err := h.orderService.CreateOrder(c.Request.Context(), req)
	if err != nil {
		// Handle specific errors using errors.Is() to check wrapped errors
		if errors.Is(err, service.ErrInsufficientStock) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if errors.Is(err, service.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		if errors.Is(err, service.ErrInvalidQuantity) || errors.Is(err, service.ErrInvalidDiscount) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Generic error - log the actual error for debugging but return user-friendly message
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create order",
			"details": err.Error(), // Include details for debugging (remove in production if needed)
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

