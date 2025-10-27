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
<script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
<style>
  body { font-family:sans-serif; background:#f6f8fa; padding:2rem; }
  h1 { color:#333; }
  canvas { max-width:700px; margin-top:2rem; }
  .metrics { display:flex; gap:2rem; flex-wrap:wrap; }
  .card { background:white; padding:1rem 1.5rem; border-radius:10px;
          box-shadow:0 2px 5px rgba(0,0,0,0.1); min-width:160px; }
  .metric { font-size:1.4rem; color:#007acc; }
</style>
</head>
<body>
<h1>📈 Notes API Live Metrics</h1>

<div class="metrics">
  <div class="card"><strong>Total Requests:</strong> <span id="requests" class="metric">0</span></div>
  <div class="card"><strong>Avg Duration:</strong> <span id="duration" class="metric">0</span> s</div>
  <div class="card"><strong>Total Errors:</strong> <span id="errors" class="metric">0</span></div>
  <div class="card"><strong>Memory:</strong> <span id="memory" class="metric">0</span> MB</div>
  <div class="card"><strong>Uptime:</strong> <span id="uptime" class="metric">0</span> s</div>
</div>

<canvas id="reqChart"></canvas>
<canvas id="latChart"></canvas>

<script>
let labels=[], reqData=[], durData=[];
const reqCtx=document.getElementById('reqChart'), latCtx=document.getElementById('latChart');

const reqChart=new Chart(reqCtx,{type:'line',
  data:{labels, datasets:[{label:'Total Requests',data:reqData, borderColor:'#007acc',fill:false,tension:.2}]},
  options:{scales:{x:{title:{text:'Time',display:true}},y:{title:{text:'Requests',display:true},beginAtZero:true}}}
});

const latChart=new Chart(latCtx,{type:'line',
  data:{labels, datasets:[{label:'Avg Duration (s)',data:durData, borderColor:'#e07a5f',fill:false,tension:.2}]},
  options:{scales:{x:{title:{text:'Time',display:true}},y:{title:{text:'Seconds',display:true},beginAtZero:true}}}
});

async function fetchMetrics(){
  const r=await fetch('/metrics'); const d=await r.json();
  document.getElementById('requests').textContent=d.total_requests;
  document.getElementById('duration').textContent=d.avg_duration.toFixed(3);
  document.getElementById('errors').textContent=d.total_errors;
  document.getElementById('memory').textContent=d.memory_mb.toFixed(2);
  document.getElementById('uptime').textContent=d.uptime_seconds;

  const t=new Date().toLocaleTimeString();
  labels.push(t); reqData.push(d.total_requests); durData.push(d.avg_duration);
  if(labels.length>20){labels.shift(); reqData.shift(); durData.shift();}
  reqChart.update(); latChart.update();
}
setInterval(fetchMetrics,3000); window.onload=fetchMetrics;
</script>
</body>
</html>
`))

// RenderDashboardHTML renders the dashboard to a sting for Fiber
func RenderDashboardHTML() (string, error) {
	var buf bytes.Buffer
	err := dashboardTemplate.Execute(&buf, nil)
	return buf.String(), err
}
