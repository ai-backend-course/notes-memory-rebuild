package main

import (
	"context"
	"notes-memory-rebuild/database"
	"notes-memory-rebuild/handlers"
	"notes-memory-rebuild/internal/dashboard"
	metricsPkg "notes-memory-rebuild/internal/metrics"
	"notes-memory-rebuild/middleware"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Load env vars (local or Docker)
	_ = godotenv.Load()

	//  Connect to the database
	database.Connect()

	// Configure Zerolog for better readability
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Create the web app (router + server)
	app := fiber.New() // Returns a Fiber app instance

	// Register the middleware early
	app.Use(middleware.ErrorHandler)
	app.Use(middleware.MetricsMiddleware)
	app.Use(middleware.RequestTimer)

	//  Register a route:
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "✅ Notes API is live and connected to the database!",
		})
	})
	app.Get("/health", handlers.Health) // When someone GETs /health, call handlers.Health
	app.Get("/metrics", middleware.MetricsHandler)
	app.Get("/metrics/history", func(c *fiber.Ctx) error {
		history := metricsPkg.ReadAll()
		return c.JSON(history)
	})
	app.Get("/metrics/export", func(c *fiber.Ctx) error {
		if err := metricsPkg.ExportCSV(); err != nil {
			return c.Status(500).SendString("Failed to export CSV")
		}
		return c.Download("metrics_history.csv")
	})
	app.Post("/notes", handlers.CreateNote)       //When a client sends a POST request to /notes, run the CreateNote function from handlers
	app.Get("/notes", handlers.GetNotes)          //When a client sends a GET request to /notes, this will retrieve all notes.
	app.Put("/notes/:id", handlers.UpdateNote)    // :id is a path parameter-- to capture a specific note's ID
	app.Delete("/notes/:id", handlers.DeleteNote) // :id is a path parameter --- to delete the specific note by ID
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html; charset=utf-8")
		html, err := dashboard.RenderDashboardHTML()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to render dashboard")
		}
		return c.SendString(html)
	})

	// Figure out which port to use - Read variables safely
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // local default so it "just works" on your machine
	}
	// Configure Zerolog based on environment
	env := os.Getenv("APP_ENV") // e.g. "development" or "production"

	if env == "production" {
		// JSON logs for cloud dashboards
		zerolog.TimeFieldFormat = time.RFC3339
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	} else {
		// Switch to human-readable console output during local dev
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
			NoColor:    false, // enable color
		})
	}

	// Start server in goroutine
	go func() {
		log.Info().
			Str("port", port).
			Msg("🚀 Starting server")

		if err := app.Listen(":" + port); err != nil {
			log.Fatal().Err(err).Msg("❌ Server failed")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // Block until interrupt (Ctrl + C or Docker Stop)

	log.Info().Msg("🧹 Gracefully shutting down server...")

	// Allow 5 seconds to finish requests
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Error().Err(err).Msg("Shutdown error")
	}

	log.Info().Msg("✅ Server stopped cleanly")

}
