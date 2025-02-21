# Redis Monitor

Redis Monitor is a performance monitoring tool designed to track Redis memory usage, CPU load, and slow queries. It provides alerts when specified thresholds are exceeded and integrates with `Telex` via a webhook.

## Features
- Monitors Redis memory usage, CPU usage, and slow queries.
- Configurable alert thresholds.
- Webhook integration to return monitoring results.
- Cron-based scheduling for periodic checks.
- JSON API response with detailed output information.

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
│── configs/               # Configuration files (if using YAML, JSON, or ENV)
│   ├── config.yaml        # App config (Redis URL, webhook URL, interval settings)
│── go.mod                 # Go module file
│── go.sum                 # Dependencies checksum
│── README.md              # Documentation

```

## Installation
### Prerequisites
Ensure you have the following installed:
- Go (latest stable version)
- Redis server
- Redis CLI

### Clone the Repository
```sh
git clone https://github.com/telexintegrations/redis-monitoring-telex-integration.git
cd redis-monitoring-telex-integration
```

### Build and Run
```sh
go build -o redis-monitor
./redis-monitor
```

## Configuration
The monitoring tool is configured via JSON payloads. Below is an example configuration:

```json
{
  "channel_id": "monitoring_channel",
  "settings": [
    { "label": "redis_url", "type": "text", "default": "redis://localhost:6379" },
    { "label": "memory_threshold", "type": "number", "default": 80 },
    { "label": "cpu_threshold", "type": "number", "default": 80 },
    { "label": "slow_query_limit", "type": "number", "default": 5 },
    { "label": "interval", "type": "text", "default": "0 * * * *" }
  ]
}
```

## API Endpoints
### Integration Handler
**Endpoint:** `/integration.json`
**Method:** `GET`

Returns integration details:
```json
{
  "data": {
    "descriptions": {
      "app_name": "Redis Monitor",
      "app_description": "Monitors Redis memory, CPU, and slow queries",
      "app_logo": "https://i.imgur.com/7JQ7JEX.png",
      "app_url": "http://localhost:8000"
    },
    "integration_category": "Performance Monitoring",
    "integration_type": "interval",
    "settings": [ ... ]
  }
}
```

## Running the Monitor Task
To manually trigger the monitoring task:
```sh
curl -X POST http://localhost:8000/tick -H "Content-Type: application/json" -d '{"channel_id": "monitoring_channel"}'
```

## Scheduling with Crontab
To run the monitor every minute, add the following entry to crontab:
```sh
* * * * * /path/to/redis-monitor
```

## License
This project is licensed under the MIT License.




