# Day 21 - Automatic Migrations & Entrypoint Script
- entrypoint.sh waits for DB, runs migrations, then starts API,
- pg_isready ensures DB is up before querying.
- psql -f executes SQL file.
- ENTRYPOINT in Dockerfile runs this script automatically.
- Guarantees DB schema exists every startup.
- docker compose up now brings entire backend online, no manual SQL.