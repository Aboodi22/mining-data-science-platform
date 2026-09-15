import pandas as pd
import json
import sqlite3
from datetime import datetime
import os

os.makedirs("output", exist_ok=True)

def load_data():
    try:
        conn = sqlite3.connect("../backend-go/mining_data.db")
        df = pd.read_sql("SELECT * FROM mining_data", conn)
        conn.close()
        print(f"✅ Loaded {len(df)} records from database")
        print(f"📋 Columns found: {list(df.columns)}")
        return df
    except Exception as e:
        print(f"⚠ Error: {e}")
        return pd.DataFrame()

if __name__ == "__main__":
    df = load_data()
    if df.empty:
        print("⚠ No data — make sure Go server is running and sample data loaded!")
        exit()
    
    report = {
        "generated_at": datetime.utcnow().isoformat(),
        "overview": {
            "total_records": int(len(df)),
            "total_ore_tons": round(df["ore_extracted"].sum(), 2),
            "avg_grade_pct": round(df["grade"].mean(), 2),
            "avg_safety_score": round(df["safety_score"].mean(), 2),
        }
    }

    with open("output/summary_report.json", "w") as f:
        json.dump(report, f, indent=2)
    
    print("✅ Analysis complete → output/summary_report.json")
    print(json.dumps(report, indent=2))