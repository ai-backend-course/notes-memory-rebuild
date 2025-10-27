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
  :root {
    --bg: #f6f8fa;
    --text: #222;
    --card-bg: #fff;
  }
  body.dark {
    --bg: #1e1e1e;
    --text: #ddd;
    --card-bg: #2c2c2c;
  }
  body {
    font-family: sans-serif;
    background: var(--bg);
    color: var(--text);
    padding: 2rem;
    transition: background .3s, color .3s;
  }
  h1 { color: var(--text); }
  canvas { max-width: 700px; margin-top: 2rem; }
  .metrics { display: flex; gap: 1.5rem; flex-wrap: wrap; }
  .card {
    background: var(--card-bg);
    padding: 1rem 1.5rem;
    border-radius: 10px;
    box-shadow: 0 2px 5px rgba(0,0,0,0.1);
    min-width: 150px;
    transition: background .3s;
  }
  .metric { font-size: 1.4rem; }
  #themeToggle { margin-top: 1rem; cursor: pointer; padding:.5rem 1rem; border:none; border-radius:6px; background:#007acc; color:white; }
</style>
</head>
<body>
<h1>⚙️ Notes API Advanced Metrics</h1>
<button id="themeToggle">🌙 Toggle Theme</button>

<div class="metrics">
  <div class="card"><strong>Total Requests:</strong> <span id="requests" class="metric">0</span></div>
  <div class="card"><strong>Avg Duration:</strong> <span id="duration" class="metric">0</span> s</div>
  <div class="card"><strong>Total Errors:</strong> <span id="errors" class="metric">0</span></div>
  <div class="card"><strong>Error Rate:</strong> <span id="errRate" class="metric">0</span>%</div>
  <div class="card"><strong>Memory:</strong> <span id="memory" class="metric">0</span> MB</div>
  <div class="card"><strong>Uptime:</strong> <span id="uptime" class="metric">0</span> s</div>
</div>

<canvas id="reqChart"></canvas>
<canvas id="latChart"></canvas>
<canvas id="errChart"></canvas>
<canvas id="upChart"></canvas>

<script>
let labels=[], reqData=[], durData=[], errData=[], upData=[];
const reqCtx=document.getElementById('reqChart'),
      latCtx=document.getElementById('latChart'),
      errCtx=document.getElementById('errChart'),
      upCtx=document.getElementById('upChart');

const makeChart=(ctx,label,color,yLabel)=>new Chart(ctx,{type:'line',
  data:{labels,datasets:[{label,data:[],borderColor:color,fill:false,tension:.2}]},
  options:{scales:{x:{title:{text:'Time',display:true}},
                   y:{title:{text:yLabel,display:true},beginAtZero:true}}}});

const reqChart=makeChart(reqCtx,'Total Requests','#007acc','Requests');
const latChart=makeChart(latCtx,'Avg Duration (s)','#e07a5f','Seconds');
const errChart=makeChart(errCtx,'Error Rate (%)','#d9534f','% Errors');
const upChart=makeChart(upCtx,'Uptime (s)','#28a745','Seconds');

async function fetchMetrics(){
  const r=await fetch('/metrics'); const d=await r.json();
  const errRate=d.total_requests>0?((d.total_errors/d.total_requests)*100):0;

  // Update numeric cards
  document.getElementById('requests').textContent=d.total_requests;
  document.getElementById('duration').textContent=d.avg_duration.toFixed(3);
  document.getElementById('errors').textContent=d.total_errors;
  document.getElementById('errRate').textContent=errRate.toFixed(2);
  document.getElementById('memory').textContent=d.memory_mb.toFixed(2);
  document.getElementById('uptime').textContent=d.uptime_seconds;

  // Update graphs
  const t=new Date().toLocaleTimeString();
  labels.push(t); reqData.push(d.total_requests);
  durData.push(d.avg_duration); errData.push(errRate); upData.push(d.uptime_seconds);

  [reqData,durData,errData,upData].forEach(arr=>{ if(arr.length>20) arr.shift(); });
  if(labels.length>20) labels.shift();

  reqChart.data.labels=latChart.data.labels=errChart.data.labels=upChart.data.labels=labels;
  [reqChart,latChart,errChart,upChart].forEach((chart,i)=>{
    chart.data.datasets[0].data=[reqData,durData,errData,upData][i];
    chart.update();
  });

  // Color feedback
  const errCard=document.getElementById('errRate').parentElement;
  errCard.style.color=errRate<2?'#28a745':errRate<5?'#f0ad4e':'#d9534f';
}
setInterval(fetchMetrics,3000); window.onload=fetchMetrics;

// Theme toggle
const toggle=document.getElementById('themeToggle');
toggle.onclick=()=>{ document.body.classList.toggle('dark');
  toggle.textContent=document.body.classList.contains('dark')?'☀️ Light Mode':'🌙 Dark Mode';
};
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
