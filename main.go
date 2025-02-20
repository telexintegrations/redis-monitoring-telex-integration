package main

import (
	"context"
	"log"
	"os"
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

	// Get Redis host from environment variable with a default fallback
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost:6379"
		log.Println("⚠️ Using default Redis host:", redisHost)
	} else {
		log.Println("✅ Using Redis host from env:", redisHost)
	}

	// Initialize Redis monitor with the configured host
	redisMon := monitor.NewRedisMonitor(redisHost)

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
