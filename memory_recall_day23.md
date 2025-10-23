# Day 23 - Environment Variables & Logging
- .env holds config; keep it gitignored
- godotenv.Load() reads .env locally; Docker passes env vars automatically.
- os.Getenv() fetches values safely.
- Structured logs use key=value format for prodution.
- LOG_LEVEL controls verbosity for future logging packages. 