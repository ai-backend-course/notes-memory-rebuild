package main

import (
	"context"
	"notes-memory-rebuild/database"
	"notes-memory-rebuild/handlers"
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
	// Configure zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Load variables from .env
	if err := godotenv.Load(); err == nil {
		log.Info().Msg("✅ Loaded environment variables from .env")
	} else {
		log.Info().Msg("ℹ️ Using Docker or system environment variables")
	}

	// Figure out which port to use - Read variables safely
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // local default so it "just works" on your machine
	}

	// Create the web app (router + server)
	app := fiber.New() // Returns a Fiber app instance

	//  Connect to the database
	database.Connect()

	//  Register a route:
	app.Get("/health", handlers.Health)           // When someone GETs /health, call handlers.Health
	app.Post("/notes", handlers.CreateNote)       //When a client sends a POST request to /notes, run the CreateNote function from handlers
	app.Get("/notes", handlers.GetNotes)          //When a client sends a GET request to /notes, this will retrieve all notes.
	app.Put("/notes/:id", handlers.UpdateNote)    // :id is a path parameter-- to capture a specific note's ID
	app.Delete("/notes/:id", handlers.DeleteNote) // :id is a path parameter --- to delete the specific note by ID

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
