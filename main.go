package main

import (
	"context"
	"fmt"
	"log"

	"redis-monitor/internal/monitor"
)

func main() {
	mon := monitor.NewRedisMonitor("localhost:6379")
	ctx := context.Background()

	memUsage, err := mon.GetMemoryUsage(ctx)
	if err != nil {
		log.Fatalf("Error fetching memory usage: %v", err)
	}
	fmt.Println("Memory Usage:\n", memUsage)
}
