# Day 17 -- UpdateNote Handler

- PUT /notes/:id edits an existing note.
- c.Params("id) extracts ID from route.
- BodyParser() reads JSON body into struct.
- Validate title/content before query.
- UPDATE ... RETURNING ... modifies + fetches row in one SQL call.
- QueryRow().Scan() fills Note struct.
- Returns JSON of updated note.
- 404 error if note ID not found. 