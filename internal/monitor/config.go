package monitor

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	RedisURL       string        `yaml:"redis_url"`
	CheckInterval  time.Duration `yaml:"check_interval"`
	CPUThreshold   float64       `yaml:"cpu_threshold"`
	MemoryThreshold int64        `yaml:"memory_threshold"`
	SlowQueryLimit int           `yaml:"slow_query_limit"`
	WebhookURL     string        `yaml:"telex_webhook_url"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	if envURL := os.Getenv("REDIS_URL"); envURL != "" {
		config.RedisURL = envURL
	}
	if envWebhook := os.Getenv("TELEX_WEBHOOK_URL"); envWebhook != "" {
		config.WebhookURL = envWebhook
	}

	log.Println("Configuration loaded successfully")
	return &config, nil
}
