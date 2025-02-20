package main

import (
	"context"
	"log"
	"time"

	"github.com/BerylCAtieno/redis-monitor/internal/monitor"
	"github.com/BerylCAtieno/redis-monitor/internal/notifier"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	ctx := context.Background()

	// Optional: Add a timeout to prevent long-running operations
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel() // Ensure cleanup

	redisMon := monitor.NewRedisMonitor("localhost:6379")

	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		memUsage, err := redisMon.GetMemoryUsage()
		if err != nil {
			log.Println("❌ Error fetching memory:", err)
			continue
		}
	
		slowLogs, err := redisMon.GetSlowQueryCount(ctx)
		if err != nil {
			log.Println("❌ Error fetching slow log count:", err)
			continue
		}
	
		cpuUsage, err := redisMon.GetCPUUsage()
		if err != nil {
			log.Println("❌ Error fetching CPU usage:", err)
			continue
		}
	
		alertMessage := monitor.ShouldSendAlert(memUsage, slowLogs, cpuUsage)
		if alertMessage != "" {
			notifier.SendAlert(alertMessage)
		} else {
			log.Println("✅ No alert needed.")
		}
	}
}