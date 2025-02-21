package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/BerylCAtieno/redis-monitor/internal/monitor"
)

// Integration metadata handler
func IntegrationHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"data": map[string]interface{}{
			"date": map[string]string{
				"created_at": time.Now().Format("2006-01-02"),
				"updated_at": time.Now().Format("2006-01-02"),
			},
			"descriptions": map[string]string{
				"app_name":        "Redis Monitor",
				"app_description": "Monitors Redis memory, CPU, and slow queries",
				"app_logo": "https://i.imgur.com/7JQ7JEX.png",
				"app_url":  "http://localhost:8000", //figure this out
				"background_color": "#ffffff",
			},
			"integration_category": "Performance Monitoring",
			"integration_type": "interval",
			"is_active": false, // figure this out
			"output": []map[string]interface{}{
				{"label": "redis_monitor", "value": true},
			},
			"key_features": []string{
				"Monitors Redis memory usage.",
				"Tracks Redis CPU load.",
				"Identifies slow queries.",
				"Configurable thresholds for alerts.",
			},
			"permissions": map[string]interface{}{
				"monitoring_user": map[string]interface{}{
					"always_online": true,
					"display_name":  "Performance Monitor",
				},
			},
			"settings": []map[string]interface{}{
				{"label": "redis_url", "type": "text", "required": true, "default": "redis://localhost:6379"},
				{"label": "memory_threshold", "type": "number", "required": true, "default": 80},
				{"label": "cpu_threshold", "type": "number", "required": true, "default": 80},
				{"label": "slow_query_limit", "type": "number", "required": true, "default": 5},
				{"label": "interval", "type": "text", "required": true, "default": "0 * * * *"},
			},
			"tick_url": "http://localhost:8000" + "/tick", //figure this out (URL for subscribing to Telex's clock)
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// MonitorPayload represents the payload sent to /tick
type MonitorPayload struct {
	ChannelID string `json:"channel_id"`
	ReturnURL string `json:"return_url"`
	Settings  []struct {
		Label   string      `json:"label"`
		Type    string      `json:"type"`
		Default interface{} `json:"default"`
	} `json:"settings"`
}

// TickHandler processes monitoring requests
func TickHandler(w http.ResponseWriter, r *http.Request) {
	var payload monitor.MonitorPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, `{"error": "Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	go monitor.RunMonitorTask(payload)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"status": "accepted"}`))
}
