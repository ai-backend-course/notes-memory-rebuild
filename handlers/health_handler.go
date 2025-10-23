package handlers

import "github.com/gofiber/fiber/v2"

// Health is a handler function.
// Fiber calls it when a request hits the /health route.
func Health(c *fiber.Ctx) error {
	// c is the request/response context.
	// SendString writes "OK" into the HTTP response body.
	return c.SendString("OK")
}
