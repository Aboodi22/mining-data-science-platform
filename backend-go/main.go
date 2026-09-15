package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MiningData struct {
	ID           uint    `gorm:"primarykey" json:"id"`
	SiteName     string  `json:"site_name"`
	Date         string  `json:"date"`
	Shift        string  `json:"shift"`
	EquipmentID  string  `json:"equipment_id"`
	OreExtracted float64 `json:"ore_extracted"`
	Grade        float64 `json:"grade"`
	EnergyUsed   float64 `json:"energy_used"`
	WaterUsed    float64 `json:"water_used"`
	Downtime     float64 `json:"downtime"`
	SafetyScore  float64 `json:"safety_score"`
}

var DB *gorm.DB

func loadSampleData() {
	var count int64
	DB.Model(&MiningData{}).Count(&count)
	if count > 0 {
		log.Println("✅ Sample data already loaded")
		return
	}

	sampleJSON := `[
		{"site_name":"North Mine","date":"2026-09-10","shift":"Morning","equipment_id":"EXCAV-01","ore_extracted":1250.5,"grade":4.85,"energy_used":245.3,"water_used":180.2,"downtime":2.5,"safety_score":92.3},
		{"site_name":"North Mine","date":"2026-09-10","shift":"Afternoon","equipment_id":"EXCAV-01","ore_extracted":1180.0,"grade":4.62,"energy_used":230.1,"water_used":175.0,"downtime":0.5,"safety_score":95.0},
		{"site_name":"North Mine","date":"2026-09-11","shift":"Morning","equipment_id":"EXCAV-02","ore_extracted":1420.8,"grade":5.10,"energy_used":270.5,"water_used":192.4,"downtime":1.0,"safety_score":88.7},
		{"site_name":"South Pit","date":"2026-09-11","shift":"Morning","equipment_id":"DRILL-05","ore_extracted":980.3,"grade":3.95,"energy_used":195.2,"water_used":140.8,"downtime":4.2,"safety_score":91.2},
		{"site_name":"South Pit","date":"2026-09-11","shift":"Afternoon","equipment_id":"DRILL-05","ore_extracted":1050.0,"grade":4.10,"energy_used":205.7,"water_used":152.0,"downtime":1.5,"safety_score":93.5},
		{"site_name":"North Mine","date":"2026-09-12","shift":"Morning","equipment_id":"EXCAV-01","ore_extracted":1310.2,"grade":4.92,"energy_used":252.0,"water_used":185.5,"downtime":3.0,"safety_score":89.0},
		{"site_name":"East Quarry","date":"2026-09-12","shift":"Night","equipment_id":"TRUCK-12","ore_extracted":760.5,"grade":6.20,"energy_used":165.8,"water_used":98.3,"downtime":0.0,"safety_score":97.2},
		{"site_name":"East Quarry","date":"2026-09-13","shift":"Morning","equipment_id":"TRUCK-12","ore_extracted":840.0,"grade":6.05,"energy_used":178.4,"water_used":105.0,"downtime":1.8,"safety_score":94.8}
	]`

	var records []MiningData
	json.Unmarshal([]byte(sampleJSON), &records)
	for _, r := range records {
		DB.Create(&r)
	}
	log.Printf("✅ Loaded %d sample records\n", len(records))
}

func initDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("mining_data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	DB.AutoMigrate(&MiningData{})
	loadSampleData()
	log.Println("✅ Database ready with sample data")
}

func main() {
	initDB()
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Next()
	})

	r.POST("/api/data", func(c *gin.Context) {
		var d MiningData
		if err := c.ShouldBindJSON(&d); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		DB.Create(&d)
		c.JSON(http.StatusCreated, gin.H{"status": "success", "data": d})
	})

	r.GET("/api/data", func(c *gin.Context) {
		var a []MiningData
		DB.Find(&a)
		c.JSON(http.StatusOK, a)
	})

	r.GET("/api/data/summary", func(c *gin.Context) {
		type Result struct {
			TotalOre    float64
			TotalEnergy float64
			TotalWater  float64
			TotalCount  int64
		}
		var res Result
		DB.Model(&MiningData{}).
			Select("sum(ore_extracted) as total_ore, sum(energy_used) as total_energy, sum(water_used) as total_water, count(*) as total_count").
			Scan(&res)

		c.JSON(http.StatusOK, gin.H{
			"total_records":    res.TotalCount,
			"total_ore_tons":   res.TotalOre,
			"total_energy_mwh": res.TotalEnergy,
			"total_water_m3":   res.TotalWater,
		})
	})

	r.GET("/api/analytics/refresh", func(c *gin.Context) {
		cmd := exec.Command("python", "../analytics-python/eda_pipeline.py")
		out, err := cmd.CombinedOutput()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "output": string(out)})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "completed", "logs": string(out)})
	})

	log.Println("🚀 Server running on http://localhost:8080")
	r.Run(":8080")
}
