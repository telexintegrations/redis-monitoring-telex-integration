# hngx-stage3-redis-monitoring-integration
Redis monitoring integration

## Directory Structure
```
redis-monitor-telex/
│── cmd/                   # Entry point for the application
│   ├── main.go            # Starts the HTTP server and loads configurations
│── internal/              # Internal package (not exposed outside the module)
│   ├── monitor/           # Redis monitoring logic
│   │   ├── monitor.go     # Collects Redis metrics (memory, CPU, slow queries)
│   │   ├── alert.go       # Logic for sending alerts to Telex
│   │   ├── config.go      # Configuration settings (intervals, Redis connection)
│   ├── server/            # HTTP server logic
│   │   ├── handlers.go    # HTTP handlers for Telex integration (tick_url, target_url)
│   │   ├── server.go      # Starts the HTTP server
│── pkg/                   # Public helper utilities (if needed)
│   ├── httpclient/        # Utility functions for making HTTP requests to Telex
│   ├── logger/            # Custom logging utilities
│── configs/               # Configuration files (if using YAML, JSON, or ENV)
│   ├── config.yaml        # App config (Redis URL, webhook URL, interval settings)
│── go.mod                 # Go module file
│── go.sum                 # Dependencies checksum
│── README.md              # Documentation

```


