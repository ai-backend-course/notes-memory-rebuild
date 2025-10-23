# Notes API - AI Backend layer (Go + Postgres)

This project is part of my **AI Backend Course**, where I'm mastering Go, Postgres, and production backend design by rebuilding everything with AI guidance and deploying real services.

--

## Overview
The **Notes API** is a foundational project that demonstrates: 
- Go + Fiber for REST API development
- Postgres + pgxpool for database connection pooling
- Context-based query timeouts for stability
- CRUD operations (Create, Read, Update, Delete)
- JSON request/response handling
- Error handling and validation

It serves as the **core microservice** in my AI Backend portfolio before extending into AI-integrated APIs (e.g., RAG Engine, AI Chat Service).

---

## Tech Stack
Component       |           Technology
**Language**                Go (1.25)
**Framework**               Fiber v2
**Database**                PostgreSQL
**Driver**                  pgx/v5 (connection pool)
**Testing**                 `testing` + `testing/assert`
**Deployment**              Docker + Railway / AWS (coming soon)

--


## Setup & Run
### 1. Clone the repository
```bash
https://github.com/ai-backend-course/notes-memory-rebuild.git
cd notes-memory-rebuild