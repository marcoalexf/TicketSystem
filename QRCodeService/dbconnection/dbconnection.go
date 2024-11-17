package dbconnection

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// Database is the struct that holds the connection pool
type Database struct {
	dsn  string
	pool *pgxpool.Pool
}

// NewDatabase initializes a new Database instance with a given DSN
func NewDatabase() *Database {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Could not load .env file, proceeding with system environment variables.")
	}

	return &Database{}
}

func (db *Database) Open() (*pgxpool.Pool, error) {
	if db.pool != nil {
		return db.pool, nil // Return existing pool if already opened
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set in the environment")
	}

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Test the connection
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	db.pool = pool
	return db.pool, nil
}

func (db *Database) Close() {
	if db.pool != nil {
		db.pool.Close()
	}
}
