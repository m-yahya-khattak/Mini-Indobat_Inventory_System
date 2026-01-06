package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB holds the database connection pool
var DB *pgxpool.Pool

// InitDB initializes the database connection pool
func InitDB(cfg *Config) error {
	config, err := pgxpool.ParseConfig(cfg.GetDBConnectionString())
	if err != nil {
		return fmt.Errorf("failed to parse database config: %w", err)
	}

	config.MaxConns = 25
	config.MinConns = 5
	config.MaxConnLifetime = 5 * time.Minute
	config.MaxConnIdleTime = 1 * time.Minute

	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return fmt.Errorf("failed to create connection pool: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := DB.Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection established successfully")
	return nil
}

// CloseDB closes the database connection pool
func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}

