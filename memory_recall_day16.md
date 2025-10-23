# Day 16 - GetNotes Handler

- GET /notes fetches all notes ordered by newest first.
- database.Pool.Query() gets multiple rows.
- rows.Next() iterates; rows.Scan() maps each column. 
- append(notes, n) builds slice of notes.
- c.JSON(notes) encodes slice to JSON response.
- context.WithTimeout() limits DB time per request. 


