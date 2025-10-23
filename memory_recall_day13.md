# Day 13 Recall

- database.Connect() reads DATABASE_URL, creates a connection pool, and pings the Database.
- Uses context.WithTimeout() to avoid waiting forever.
- pool.Ping() ensures the database is alive before the app starts.
- The global variable Pool is used by handlers to run queries.
- Called from main.go before starting the server. 

# Day 13 Phase 2 Notes
- Pool.QueryRow(ctx, "SELECT NOW()") sends SQL to Postgres.
- .Scan(&variable) copies column data into Go variable.
- QueryRow is for single-row results; Query is for multiple.
- context.WithTimeout prevents long waits.
- go test ./database runs the test automatically. 

