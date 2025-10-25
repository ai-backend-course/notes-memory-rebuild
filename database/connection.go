package database

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

var Pool *pgxpool.Pool

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal().Msg("❌ DATABASE_URL is not set")
	}

	var pool *pgxpool.Pool
	var err error

	for i := 1; i <= 5; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		pool, err = pgxpool.New(ctx, dsn)
		if err == nil {
			if err = pool.Ping(ctx); err == nil {
				Pool = pool
				log.Info().Msg("✅ Database connected successfully!")
				return
			}
		}

		log.Warn().Msgf("Retrying DB connection... attempt %d/5", i)
		time.Sleep(3 * time.Second)
	}

	log.Fatal().Err(err).Msg("❌ Failed to connect to database after retries")
}
