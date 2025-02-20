package monitor

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Define global variables for thresholds
var (
	MemoryThreshold  int64
	SlowLogThreshold int
	CPUThreshold     float64
)

// Load environment variables
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	MemoryThreshold = getEnvAsInt64("MEMORY_THRESHOLD", 100*1024*1024)
	SlowLogThreshold = getEnvAsInt("SLOW_LOG_THRESHOLD", 10)
	CPUThreshold = getEnvAsFloat("CPU_THRESHOLD", 80.0)

	log.Println("✅ Thresholds Loaded - Memory:", MemoryThreshold, "Slow Logs:", SlowLogThreshold, "CPU:", CPUThreshold)
}

// Helper functions to parse environment variables
func getEnvAsInt64(key string, defaultValue int64) int64 {
	val, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	parsed, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		log.Printf("⚠️ Invalid value for %s, using default: %d\n", key, defaultValue)
		return defaultValue
	}
	return parsed
}

func getEnvAsInt(key string, defaultValue int) int {
	val, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	parsed, err := strconv.Atoi(val)
	if err != nil {
		log.Printf("⚠️ Invalid value for %s, using default: %d\n", key, defaultValue)
		return defaultValue
	}
	return parsed
}

func getEnvAsFloat(key string, defaultValue float64) float64 {
	val, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	parsed, err := strconv.ParseFloat(val, 64)
	if err != nil {
		log.Printf("⚠️ Invalid value for %s, using default: %.2f\n", key, defaultValue)
		return defaultValue
	}
	return parsed
}

// Function to check if an alert should be sent
func ShouldSendAlert(memUsage int64, slowLogs int, cpuUsage float64) string {
	if memUsage > MemoryThreshold {
		return fmt.Sprintf("🚨 High Memory Usage: %s", formatBytes(memUsage))
	} else if slowLogs > SlowLogThreshold {
		return fmt.Sprintf("⚠️ High Slow Query Count: %d", slowLogs)
	} else if cpuUsage > CPUThreshold {
		return fmt.Sprintf("🔥 High CPU Usage: %.2f%%", cpuUsage)
	}
	return ""
}

// Format bytes into a readable format
func formatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}
