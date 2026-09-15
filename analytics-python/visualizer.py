import plotly.express as px
import pandas as pd
import sqlite3
import os

os.makedirs("../web-ui/charts", exist_ok=True)

try:
    conn = sqlite3.connect("../backend-go/mining_data.db")
    df = pd.read_sql("SELECT * FROM mining_data", conn)
    conn.close()
    print(f"✅ Loaded {len(df)} records for visualization")
except Exception as e:
    print(f"⚠ Error: {e}")
    print("⚠ Make sure Go server is running!")
    exit()

# Chart 1: Production by Site & Date
fig1 = px.bar(
    df, x="date", y="ore_extracted", color="site_name", barmode="group",
    title="📊 Daily Ore Extraction by Site (tons)",
    labels={"ore_extracted": "Ore (Tons)", "date": "Date", "site_name": "Site"}
)
fig1.write_html("../web-ui/charts/production.html")
print("✅ Created: production.html")

# Chart 2: Grade vs Downtime
fig2 = px.scatter(
    df, x="grade", y="downtime", color="site_name", size="energy_used",
    title="⚙ Ore Grade vs Equipment Downtime"
)
fig2.write_html("../web-ui/charts/grade.html")
print("✅ Created: grade.html")

# Chart 3: Resource Usage
fig3 = px.line(
    df, x="date", y=["energy_used", "water_used"],
    title="💡 Resource Consumption Over Time"
)
fig3.write_html("../web-ui/charts/resources.html")
print("✅ Created: resources.html")

print("\n✅ ALL CHARTS GENERATED SUCCESSFULLY! 🎉")