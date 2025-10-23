# Day 18 - DeleteNote Handler
- DELETE /notes/:id removes a note by ID.
- Pool.Exec() runs DELETE queries (no return rows).
- cmd.RowsAffected() checks if a record was deleted.
- Responds with JSON message or 404 error.
- Completes full CRUD cycle (Create, Read, Update, Delete).