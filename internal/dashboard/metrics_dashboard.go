package dashboard

import (
	"bytes"
	"html/template"
)

// MetricSnapshot mirrors your /metrics JSON
type MetricSnapshot struct {
	TotalRequests int     `json:"total_requests"`
	AvgDuration   float64 `json:"avg_duration"`
	TotalErrors   int     `json:"total_errors"`
	UptimeSeconds int64   `json:"uptime_seconds"`
	MemoryMB      float64 `json:"memory_mb"`
	StartTime     string  `json:"start_time"`
}

var dashboardTemplate = template.Must(template.New("dash").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>Notes API Metrics Dashboard</title>
<script>
async function fetch Metrics()  {
	const res  =  await fetch('/metrics');
	const data =  await  res.json();
	document.getElementById('requests').innerText = data.total_requests;
	document.getElementById('duration').innerText=data.avg_duration.toFixed(3);
	document.getElementById('errors').innerText = data.total_errors;
	document.getElementById('uptime').innerText = data.uptime_seconds + 's';
	document.getElementById('memory').innerText = data.memory_mb.toFixed(2)
 + ' MB';
}
 setInterval(fetchMetrics, 3000);
 window.onload = fetchMetrics;
 </script>
 <style>
body { font-family: sans-serif; background: #fafafa; padding: 2rem; }
h1 { color: #333; }
.card { background: white; padding: 1rem 2rem; border-radius: 10px; box-shadow: 0 2px 5px rgba(0,0,0,0.1); margin-bottom: 1rem; }
.metric { font-size: 1.5rem; color: #007acc; }
</style>
</head>
<body>
<h1>📊 Notes API Metrics</h1>
<div class="card"><strong>Total Requests:</strong> <span id="requests" class="metric">0</span></div>
<div class="card"><strong>Average Duration:</strong> <span id="duration" class="metric">0</span> s</div>
<div class="card"><strong>Total Errors:</strong> <span id="errors" class="metric">0</span></div>
<div class="card"><strong>Uptime:</strong> <span id="uptime" class="metric">0</span></div>
<div class="card"><strong>Memory Usage:</strong> <span id="memory" class="metric">0</span></div>
</body>
</html>
`))

// RenderDashboardHTML renders the dashboard to a sting for Fiber
func RenderDashboardHTML() (string, error) {
	var buf bytes.Buffer
	err := dashboardTemplate.Execute(&buf, nil)
	return buf.String(), err
}
