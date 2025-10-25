package middleware

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

var (
	mu            sync.Mutex
	totalRequests int
	totalDuration time.Duration
)

// MetricsMiddleware tracks total requests and average response time.
func MetricsMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	duration := time.Since(start)

	mu.Lock()
	totalRequests++
	totalDuration += duration
	avg := totalDuration / time.Duration(totalRequests)
	mu.Unlock()

	log.Info().
		Int("total_requests", totalRequests).
		Dur("avg_duration", avg).
		Msg("metrics update")

	return err
}
