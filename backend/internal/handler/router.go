package handler

import (
	"mini-indobat-inventory/internal/repository"
	"mini-indobat-inventory/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter configures and returns the HTTP router
func SetupRouter() *gin.Engine {
	// Initialize repositories
	productRepo := repository.NewProductRepository()
	orderRepo := repository.NewOrderRepository()

	// Initialize services
	productService := service.NewProductService(productRepo)
	orderService := service.NewOrderService(orderRepo, productRepo)

	// Initialize handlers
	productHandler := NewProductHandler(productService)
	orderHandler := NewOrderHandler(orderService)

	// Create router
	router := gin.Default()

	// Configure CORS for frontend
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API routes (matching requirements: GET /products, POST /products, POST /order)
	router.GET("/products", productHandler.GetProducts)
	router.POST("/products", productHandler.CreateProduct)
	router.POST("/order", orderHandler.CreateOrder)

	return router
}

