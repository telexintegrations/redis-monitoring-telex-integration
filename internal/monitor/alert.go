package monitor

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Alert defines the structure of an alert message.
type Alert struct {
	Timestamp   time.Time `json:"timestamp"`
	Metric      string    `json:"metric"`
	Value       any       `json:"value"`
	Description string    `json:"description"`
}

func (a *Alert) SendAlert(s string) {
	panic("unimplemented")
}

// TelexClient handles sending alerts to Telex.
type TelexClient struct {
	WebhookURL string
}

// NewTelexClient initializes a TelexClient with a webhook URL.
func NewTelexClient(webhookURL string) *TelexClient {
	return &TelexClient{WebhookURL: webhookURL}
}

// SendAlert sends an alert to the Telex webhook.
func (t *TelexClient) SendAlert(metric string, value any, description string) error {
	alert := Alert{
		Timestamp:   time.Now(),
		Metric:      metric,
		Value:       value,
		Description: description,
	}

	payload, err := json.Marshal(alert)
	if err != nil {
		return err
	}

	resp, err := http.Post(t.WebhookURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to send alert: %s\n", resp.Status)
	}

	return nil
}
