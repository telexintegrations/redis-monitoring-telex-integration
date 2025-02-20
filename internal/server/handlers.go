package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/BerylCAtieno/redis-monitor/internal/monitor"
)

// HealthCheckResponse represents the structure of the health check API response
type HealthCheckResponse struct {
	Status string `json:"status"`
}

// TickHandler handles periodic tick requests from Telex
func TickHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received Telex tick request")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HealthCheckResponse{Status: "ok"})
}

// AlertHandler triggers Redis monitoring and sends alerts if needed
func AlertHandler(redisMonitor *monitor.RedisMonitor, alert *monitor.Alert) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received alert trigger request")

		// Check Redis memory usage
		memUsage, err := redisMonitor.GetMemoryUsage()
		if err != nil {
			http.Error(w, "Failed to retrieve memory usage", http.StatusInternalServerError)
			return
		}

		// Check slow queries
		slowQueries, err := redisMonitor.GetSlowQueryCount(r.Context())
		if err != nil {
			http.Error(w, "Failed to retrieve slow query count", http.StatusInternalServerError)
			return
		}

		// Check CPU usage
		cpuUsage, err := redisMonitor.GetCPUUsage()
		if err != nil {
			http.Error(w, "Failed to retrieve CPU usage", http.StatusInternalServerError)
			return
		}

		// Determine if alerts should be sent
		if memUsage >= monitor.Config.MemoryThreshold {
			alert.SendAlert("Memory usage exceeded threshold!")
		}
		if slowQueries >= monitor.Config.SlowQueryLimit {
			alert.SendAlert("Slow query count exceeded threshold!")
		}
		if cpuUsage >= monitor.Config.CPUThreshold {
			alert.SendAlert("CPU usage exceeded threshold!")
		}

		// Send a success response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(HealthCheckResponse{Status: "Monitoring check completed"})
	}
}
