# ⛏️ Mining Data Science Platform

> Full-stack mining analytics platform with Go backend API, SQLite database, Python analytics pipeline, and interactive web dashboard.

---

## 📋 Overview

A complete data science solution for mining operations:
- **Go REST API** — backend & database (SQLite + GORM)
- **Python Analytics** — automated EDA, KPIs, and interactive charts
- **Web Dashboard** — real-time visualization, data entry, CSV bulk import, mine location mapping

---

## 🗂️ Project Structure
mining-data-science-platform/
├── backend-go/ # Go API Server
│ ├── main.go # REST API endpoints
│ ├── go.mod # Dependencies
│ └── mining_data.db # SQLite database (auto-created)
├── analytics-python/ # Python Analytics Pipeline
│ ├── eda_pipeline.py # Data analysis & KPIs
│ ├── visualizer.py # Interactive charts (Plotly)
│ ├── requirements.txt # Python dependencies
│ ├── output/ # Generated reports
│ └── ...
├── web-ui/ # Dashboard
│ ├── index.html # Main dashboard
│ └── charts/ # Generated HTML charts
├── data/ # Raw & processed data
├── .gitignore
└── README.md

---

## 🚀 Quick Start

### Prerequisites
- **Go 1.22+**
- **Python 3.11+**
- **GCC / MinGW-w64** (for SQLite CGO support on Windows)

### 1️⃣ Start the Go Backend
```powershell
cd backend-go
go mod tidy
$env:CGO_ENABLED="1"
go run main.go
Server runs at: http://localhost:8080
2️⃣ Run Python Analytics
cd analytics-python
pip install -r requirements.txt
python eda_pipeline.py
python visualizer.py
3️⃣ Open Dashboard
Open web-ui/index.html in your browser
📊 Features
Feature	Status
REST API with SQLite	✅
Sample mining data auto-loaded	✅ (8 records)
Manual data entry form	✅
CSV / Excel bulk import	✅
Production & resource charts	✅
Advanced analytics (shift, efficiency, safety)	✅
Mine location map	✅
Auto-refreshing KPIs	✅
📄 Data Format
CSV Import Columns:
plaintext
site_name, date, shift, equipment_id, ore_extracted, grade, energy_used, water_used, downtime, safety_score
🛠️ Built With
Go + Gin — REST API
SQLite + GORM — Database
Pandas / Plotly — Python analytics
Plotly.js — Dashboard charts
HTML/CSS/JS — Frontend
