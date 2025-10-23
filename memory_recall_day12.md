# Day 12 Recall

- main.go
    - Creates Fiber app with fiber.New()
    - Registers routes (e.g., app.Get("/health", handlers.Health))
    - Reads PORT from env (fallback is 8080)
    - app.Listen starts the HTTP server


- handlers/health_handler.go
    - Health(c *fiber.Ctx) returns "OK" (simple alive/health check)
    - Handlers take *fiber.Ctx and return error

- database/connection.go
    - Placeholder today: tomorrow it will: 
        - read DATABASE_URL
        - open pgxpool
        - Ping and log success/failure
        - expose global Pool for queries