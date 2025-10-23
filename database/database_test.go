package database

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestDatabasePing ensures the global Pool  works and can query Postgres.
func TestDatabasePing(t *testing.T) {
	// 1.) Connect to Database (opens Pool)
	Connect()

	// 2.) Use a short context for safety
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 3.) Run a simple SQL command that always works
	var currentTime time.Time
	err := Pool.QueryRow(ctx, "SELECT NOW()").Scan(&currentTime)

	// 4.) Verify
	assert.NoError(t, err)                                           // fails the test if any Database error occurred
	assert.WithinDuration(t, time.Now(), currentTime, 5*time.Second) // Ensures the time returned is close to local time
}
