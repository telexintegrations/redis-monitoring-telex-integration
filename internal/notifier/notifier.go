package notifier

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const WebhookURL = "https://ping.telex.im/v1/webhooks/01951f3f-d29a-75c2-acb9-d8b83553d2a5"

func SendAlert(message string) {
	data := map[string]string{
		"event_name": "Redis Alert",
		"message":    message,
		"status":     "warning",
		"username":   "RedisMonitor",
	}

	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", WebhookURL, bytes.NewBuffer(jsonData))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("❌ Failed to send alert:", err)
		return
	}
	defer resp.Body.Close()

	log.Println("✅ Alert sent:", resp.Status)
}
