package main

import (
	"log"
	"notes-memory-rebuild/database"
	"notes-memory-rebuild/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load variables from .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, using system env vars")
	}

	// 1) Create the web app (router + server)
	app := fiber.New() // Returns a Fiber app instance

	// 2.) Connect to the database
	database.Connect()

	// 3) Register a route:
	app.Get("/health", handlers.Health)           // When someone GETs /health, call handlers.Health
	app.Post("/notes", handlers.CreateNote)       //When a client sends a POST request to /notes, run the CreateNote function from handlers
	app.Get("/notes", handlers.GetNotes)          //When a client sends a GET request to /notes, this will retrieve all notes.
	app.Put("/notes/:id", handlers.UpdateNote)    // :id is a path parameter-- to capture a specific note's ID
	app.Delete("/notes/:id", handlers.DeleteNote) // :id is a path parameter --- to delete the specific note by ID

	// 4) Figure out which port to use
	// PORT is commonly supplied by hosting platforms (Render, Railway, etc.)
	// Read variables safely
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // local default so it "just works" on your machine
	}

	// 5) Start the server
	// app.Listen starts listening for HTTP requests on :8080 (or whatever PORT is).
	// If it fails (port in use, permissions, etc.), log.Fatal will print the error and exit.
	log.Printf("🚀 Server running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
