package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// RequestTimer measures how long each request takes and logs details
func RequestTimer(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	duration := time.Since(start)

	// Log structured metrics

	status := c.Response().StatusCode()

	// Base event with shared fields
	event := log.With().
		Str("method", c.Method()).
		Str("path", c.Path()).
		Int("status", c.Response().StatusCode()).
		Dur("duration", duration).
		Logger()

	switch {
	case status >= 500:
		event.Error().Msg("server error")
	case status >= 400:
		event.Warn().Msg("client issue")
	default:
		event.Info().Msg("request completed")
	}

	return err
}
