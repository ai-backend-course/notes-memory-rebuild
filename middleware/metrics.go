package middleware

import (
	"runtime"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// MetricsData holds live stats accessible from the /metrics endpoint
type MetricsData struct {
	TotalRequests int     `json:"total_requests"`
	AvgDuration   float64 `json:"avg_duration"`
	TotalErrors   int     `json:"total_errors"`
	StartTime     string  `json:"start_time"`
	UptimeSeconds int64   `json:"uptime_seconds"`
	MemoryMB      float64 `json:"memory_mb"`
}

var (
	metrics   MetricsData
	metricsMu sync.Mutex
	started   = time.Now()
)

// MetricsMiddleware tracks total requests and average response time.
func MetricsMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	duration := time.Since(start).Seconds()

	metricsMu.Lock()
	defer metricsMu.Unlock()

	metrics.TotalRequests++
	metrics.AvgDuration = ((metrics.AvgDuration * float64(metrics.TotalRequests-1)) + duration) / float64(metrics.TotalRequests)

	if c.Response().StatusCode() >= 400 {
		metrics.TotalErrors++
	}

	if metrics.StartTime == "" {
		metrics.StartTime = started.Format(time.RFC3339)
	}

	log.Info().
		Int("total_requests", metrics.TotalRequests).
		Float64("avg_duration", metrics.AvgDuration).
		Msg("metrics update")

	return err
}

// MetricsHandler returns current metrics as JSON
func MetricsHandler(c *fiber.Ctx) error {
	metricsMu.Lock()
	defer metricsMu.Unlock()

	// Uptime in seconds
	metrics.UptimeSeconds = int64(time.Since(started).Seconds())

	// Memory usage in MB
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	metrics.MemoryMB = float64(m.Alloc) / 1024 / 1024

	return c.JSON(metrics)
}
