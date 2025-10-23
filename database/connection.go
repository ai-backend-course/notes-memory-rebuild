package database

import (
	"context" //allows us to set timeouts for database operations
	"log"     // for logging info or errors
	"os"      // to read environment variables like DATABASE_URL
	"time"    // for timeouts

	"github.com/jackc/pgx/v5/pgxpool" // connection pool driver
	"github.com/joho/godotenv"
)

var Pool *pgxpool.Pool // shared pool of Database connections across the app

// Connect opens the global connection pool
func Connect() {
	// 0.) Load environment variables from .env
	// Try to load from current directory first, then from parent directory
	if err := godotenv.Load(); err != nil {
		// If loading from current directory fails, try parent directory (for tests)
		if err := godotenv.Load("../.env"); err != nil {
			log.Println("⚠️ No .env file found -- using system environment variables")
		}
	}
	// 1.) Get the connection string from the .env file
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL is not set in .env")
	}

	// 2.) Create a context with timeout (prevents hanging)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // ensures the context is cleaned up when function ends

	// 3.) Open a new connection pool
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("❌ Failed to create DB pool: %v", err)
	}

	// 4.) Ping the database to confirm connection
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("❌ Database ping failed: %v", err)
	}

	// 5.) Assign to global Pool and log success
	Pool = pool
	log.Println("✅ Database connected successfully!")

}
