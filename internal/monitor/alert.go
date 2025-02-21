package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// MonitorPayload represents the monitoring settings
type MonitorPayload struct {
	ChannelID string `json:"channel_id"`
	ReturnURL string `json:"return_url"`
	Settings  []struct {
		Label   string      `json:"label"`
		Type    string      `json:"type"`
		Default interface{} `json:"default"`
	} `json:"settings"`
}

// RunMonitorTask fetches Redis metrics and sends alerts
func RunMonitorTask(payload MonitorPayload) {
	ctx := context.Background()

	// Extract settings from payload
	redisURL := "redis://localhost:6379" // Default
	var memoryThreshold, cpuThreshold, slowQueryLimit int

	for _, setting := range payload.Settings {
		switch setting.Label {
		case "redis_url":
			if val, ok := setting.Default.(string); ok {
				redisURL = val
			}
		case "memory_threshold":
			if val, ok := setting.Default.(float64); ok {
				memoryThreshold = int(val)
			}
		case "cpu_threshold":
			if val, ok := setting.Default.(float64); ok {
				cpuThreshold = int(val)
			}
		case "slow_query_limit":
			if val, ok := setting.Default.(float64); ok {
				slowQueryLimit = int(val)
			}
		}
	}

	// Initialize RedisMonitor
	monitor := NewRedisMonitor(redisURL)

	// Fetch metrics
	memUsage, err := monitor.GetMemoryUsage()
	if err != nil {
		log.Println("Error getting memory usage:", err)
		return
	}

	cpuUsage, err := monitor.GetCPUUsage()
	if err != nil {
		log.Println("Error getting CPU usage:", err)
		return
	}

	slowQueries, err := monitor.GetSlowQueryCount(ctx)
	if err != nil {
		log.Println("Error getting slow query count:", err)
		return
	}

	// Check thresholds
	var alerts []string
	if memUsage >= int64(memoryThreshold)*1024*1024 {
		alerts = append(alerts, fmt.Sprintf("Memory usage exceeded: %d bytes", memUsage))
	}
	if cpuUsage >= float64(cpuThreshold) {
		alerts = append(alerts, fmt.Sprintf("CPU usage exceeded: %.2f", cpuUsage))
	}
	if slowQueries >= slowQueryLimit {
		alerts = append(alerts, fmt.Sprintf("Slow queries exceeded: %d", slowQueries))
	}

	// If no alerts, return early
	if len(alerts) == 0 {
		log.Println("No issues detected.")
		return
	}

	// Send response back
	data := map[string]interface{}{
		"message":    strings.Join(alerts, "; "),
		"username":   "Redis Monitor",
		"event_name": "Redis Health Check",
		"status":     "error",
	}

	// Send results back to the return URL
	client := &http.Client{Timeout: 10 * time.Second}
	jsonData, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", payload.ReturnURL, strings.NewReader(string(jsonData)))
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("Alert sent. Response status: %s\n", resp.Status)
}
