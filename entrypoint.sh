#!/bin/sh
set -e

echo "🐘 Waiting for Postgres..."
# Wait for DB to accept connections
until pg_isready -h db -p 5432 -U postgres > /dev/null 2>&1; do
  sleep 1
done
echo "✅ Postgres is ready!"

echo "🧩 Running migrations..."
# Apply migrations (only if table missing)
psql postgres://postgres:${POSTGRES_PASSWORD}@db:5432/${POSTGRES_DB}?sslmode=disable \
  -f migrations/000001_create_notes_table.sql

echo "🚀 Starting Notes API..."
exec ./notes-api

