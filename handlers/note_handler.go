package handlers

import (
	"context"
	"notes-memory-rebuild/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

// 1️⃣ Define the expected JSON structure
type NoteInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// 2️⃣ Define how a note looks in the Database + response
type Note struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 3️⃣ CreateNote handles POST /notes
func CreateNote(c *fiber.Ctx) error {
	var input NoteInput

	// Parse the JSON body into our struct
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON payload")
	}

	// ✅ Simple Validation
	if input.Title == "" || input.Content == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Title and Content are required")
	}

	// 4️⃣ Create a context with timeout (safety net)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 5️⃣ SQL query
	query := `
INSERT INTO notes (title, content)
VALUES ($1, $2)
RETURNING id, title, content, created_at, updated_at;
`

	// 6️⃣ Execute query and scan the returned row
	var note Note
	err := database.Pool.QueryRow(ctx, query, input.Title, input.Content).
		Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Database insert failed")
	}

	// 7️⃣ Return the new note as JSON
	return c.Status(fiber.StatusCreated).JSON(note)
}

// GetNotes handles GET /notes
func GetNotes(c *fiber.Ctx) error {
	// 1.) creates a context with Timeout (safety net)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 2.) The SQL query
	query := `
	SELECT id, title, content, created_at, updated_at
	FROM notes
	ORDER BY created_at DESC;
	`

	rows, err := database.Pool.Query(ctx, query)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch notes")
	}

	defer rows.Close()

	var notes []Note

	// Iterate over each row returned
	for rows.Next() {
		var n Note
		if err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt, &n.UpdatedAt); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to read row")

		}
		notes = append(notes, n)
	}

	return c.JSON(notes)
}

func UpdateNote(c *fiber.Ctx) error {
	// 1.)  Extract the note ID from the URL
	id := c.Params("id")

	// 2.) Parse JSON body into NoteInput struct
	var input NoteInput
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON payload")

	}

	// 3.) Validate required fields
	if input.Title == "" || input.Content == "" {
		return fiber.NewError(fiber.StatusBadRequest, "title and content are required")
	}

	// 4.) Context with timeout (safety)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 5.) SQL: update the row and return the new data
	query := `
	UPDATE notes
	SET title = $1, content = $2, updated_at = NOW()
	WHERE id = $3
	RETURNING id, title, content, created_at, updated_at;
	`

	var note Note
	err := database.Pool.QueryRow(ctx, query, input.Title, input.Content, id).
		Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Note not found or update failed")
	}
	return c.JSON(note)
}

// DeleteNote handles DELETE /notes/:id
func DeleteNote(c *fiber.Ctx) error {
	// 1. Get the ID from the URL
	id := c.Params("id")

	// 2. Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 3. Run DELETE query
	query := `DELETE FROM notes WHERE id = $1;`
	cmd, err := database.Pool.Exec(ctx, query, id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Note not found or already deleted")
	}

	// 4. Check if any row was actually deleted
	if cmd.RowsAffected() == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Note not found")
	}

	// 5. Respond with success message
	return c.JSON(fiber.Map{
		"message": "Note deleted successfully",
	})

}
