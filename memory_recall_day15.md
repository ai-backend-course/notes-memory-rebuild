# Day 15 - CreateNote Handler

- POST /notes adds a note to the DataBase.
- BodyParser() decodes JSON -> the struct.
- Validate required fields.
- context. WithTimeout() limits query duration.
- QueryRow() + Scan() runs SQL and captures the result.
- RETURNING clause instantly gives back inserted data.
- Respond with StatusCreated(201) + JSON.