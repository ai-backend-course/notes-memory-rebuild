package metrics

import (
	"encoding/json"
	"os"
	"sync"
)

type MetricSnapshot struct {
	TotalRequests int     `json:"total_requests"`
	AvgDuration   float64 `json:"avg_duration"`
	TotalErrors   int     `json:"total_errors"`
	UptimeSeconds int64   `json:"uptime_seconds"`
	MemoryMB      float64 `json:"memory_mb"`
	Timestamp     string  `json:"timestamp"`
}

var (
	historyFile = "metrics_history.json"
	mu          sync.Mutex
	maxSamples  = 200 // keep last 200 snapshots
)

// Append adds a new snapshot and trims file length
func Append(snapshot MetricSnapshot) {
	mu.Lock()
	defer mu.Unlock()

	var data []MetricSnapshot
	_ = readFile(&data)
	data = append(data, snapshot)
	if len(data) > maxSamples {
		data = data[len(data)-maxSamples:]
	}
	writeFile(data)
}

// ReadAll returns the stored history
func ReadAll() []MetricSnapshot {
	mu.Lock()
	defer mu.Unlock()
	var data []MetricSnapshot
	_ = readFile(&data)
	return data
}

func readFile(dst *[]MetricSnapshot) error {
	b, err := os.ReadFile(historyFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}

func writeFile(data []MetricSnapshot) {
	b, _ := json.MarshalIndent(data, "", "  ")
	_ = os.WriteFile(historyFile, b, 0644)
}
