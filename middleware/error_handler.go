package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// Error Handler is a global Fiber middleware.

func ErrorHandler(c *fiber.Ctx) error {
	// 1 Generate a unique request ID for every request
	reqID := uuid.New().String()
	c.Locals("request_id", reqID)

	// 2 Wrap the next handler and recover from any panic
	defer func() {
		if r := recover(); r != nil {
			log.Error().
				Str("request_id", reqID).
				Str("path", c.Path()).
				Interface("panic", r).
				Msg("Unhandled panic")
			_ = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":       "Internal Server Error",
				"request_id":  reqID,
				"description": fmt.Sprintf("%v", r),
			})
		}
	}()

	// 3 Continue processing and catch returned errors
	err := c.Next()
	if err != nil {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
		log.Error().
			Str("request_id", reqID).
			Str("path", c.Path()).
			Err(err).
			Msg("Request error")

		return c.Status(code).JSON(fiber.Map{
			"error":      err.Error(),
			"request_id": reqID,
		})
	}

	return nil
}
