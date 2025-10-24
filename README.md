# Notes API - AI Backend layer (Go + Postgres)

This project is part of my **AI Backend Course**, where I'm mastering Go, Postgres, and production backend design by rebuilding everything with AI guidance and deploying real services.

---

## Overview
The **Notes API** is a foundational project that demonstrates: 
- Go + Fiber for REST API development
- Postgres + pgxpool for database connection pooling
- Context-based query timeouts for stability
- CRUD operations (Create, Read, Update, Delete)
- JSON request/response handling
- Error handling and validation
- Testing and Docker integration (upcoming)

It serves as the **core microservice** in my AI Backend portfolio before extending into AI-integrated APIs (e.g., RAG Engine, AI Chat Service).

---

## Tech Stack
|  Component       |           Technology  |
|--------------------|-----------------------|
|**Language**         |       Go (1.25)     |
|**Framework**         |      Fiber v2      |
| **Database**         |       PostgreSQL   |
| **Driver**        |          `pgx/v5` (connection pool) |
| **Testing**       |          `testing` + `testing/assert`|
| **Deployment**      |        Docker + Railway / AWS (coming soon) |

---


## Setup & Run


### 1. Clone the repository
```bash
https://github.com/ai-backend-course/notes-memory-rebuild.git
cd notes-memory-rebuild  
```

### 2. Environment variables
Create a .env file:
```bash
DATABASE_URL=postgres://postgres:password@localhost:5432/ai_backend?sslmode=disable
```

### 3. Run the Server
```bash
go run main.go
```
Server runs on http://localhost:8080


## API Endpoints
|  Method    |      Endpoint | Description |
|--------------------|-----------------------|--------------|
|**GET**         |      /health    |  Check server health   |
|**POST**         |      /notes     |  Create a note      |
| **GET**         |      /notes  |    Fetch all notes       |
| **PUT**        |       /notes/:id |  Update an existing note  |
| **DELETE**       |     /notes/:id | Delete a note     |

## Progress Log
|  Day    |      Milestone |
|--------------------|--------------|
|**Day 11**         |      Introduced Go testing basics  |
|**Day 12**         |      Memory rebuild project skeleton   |
|**Day 13**         |      Database connection + ping test    |
| **Day 14**        |      Created `notes` table and migration |
| **Day 15**       |     Implemented a `CreateNote` handler |
|**Day 13**         |     Implemented a `GetNotes` handler   |
| **Day 14**        |      Implemented a `UpdateNote` handler |
| **Day 15**       |     Implemented a `DeleteNote` handler (Full CRUD Complete) |

---

## Logging & Graceful Shutdown
- Uses **Zerolog** for JSON-formatted logs suitable for cloud dashboards.
- Graceful shutdown via `os/signal` ensures clean container exits.
- Example:
```json
{"level":"info", "event": "startup", "port":"8080"}
```
---

## ⚠️ Centralized Error Handling
- Global Fiber middleware catches all errors and panics.
- Adds a unique `request_id` for each request.
- Returns structured JSON: 
```json
{"error":"Internal Server Error", "request_id":"<uuid>"}
```

---

## Learning Objectives
- Strengthen backend fundamentals with Go and SQL
- Repetitive Practice through the AI Backend Course
- Learn structure production-grade APIs
- Prepare for advanced AI-backend layers (vector search, embeddings, RAG APIs)

## Author 
**Jeff Ellis**  
Backend Developer | AI Backend Course  
[AI Backend Course Organization](https://github.com/ai-backend-course/notes-memory-rebuild)
