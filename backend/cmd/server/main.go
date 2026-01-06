package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mini-indobat-inventory/internal/config"
	"mini-indobat-inventory/internal/handler"
)

func main() {
	log.Println("Mini-Indobat Inventory System - Backend Server")
	log.Println("Server starting...")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := config.InitDB(cfg); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer config.CloseDB()

	router := handler.SetupRouter()

	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: router,
	}

	go func() {
		log.Printf("Server listening on port %s", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
