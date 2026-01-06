package handler

import (
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
		// Handle specific errors
		if err == service.ErrInsufficientStock {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err == service.ErrProductNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err == service.ErrInvalidQuantity || err == service.ErrInvalidDiscount {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Generic error
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create order",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

