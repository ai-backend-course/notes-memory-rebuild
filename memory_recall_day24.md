# Day 24 - Structured Logging & Graceful Shutdown
- Zerolog provides JSON logs with key=value fields.
- Structured logs are machine-readable for observability.
- Graceful shutdown catches SIGINT/SIGTERM signals.
- app.Shutdown() allows active requests to finish.
- Context timeout ensures container stops within 5s. 