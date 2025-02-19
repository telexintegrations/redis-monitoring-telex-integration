package monitor

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisMonitor struct {
	Client *redis.Client
}

func NewRedisMonitor(addr string) *RedisMonitor {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisMonitor{Client: client}
}

// GetMemoryUsage retrieves Redis memory usage in bytes.
func (r *RedisMonitor) GetMemoryUsage() (int64, error) {
	mem, err := r.Client.Info(ctx, "memory").Result()
	if err != nil {
		return 0, err
	}

	var usage int64
	fmt.Sscanf(mem, "used_memory:%d", &usage)
	return usage, nil
}

// get slow queries count 
func (r *RedisMonitor) GetSlowQueryCount(ctx context.Context) (int, error) {
	res, err := r.Client.Do(ctx, "SLOWLOG", "GET").StringSlice()
	if err != nil {
		return 0, err
	}

	// Each slow log entry consists of multiple fields, so we count logs by dividing total entries
	// Redis slow log entries have 4+ fields per log (varies by Redis version)
	const slowLogEntrySize = 4 // Change if your Redis version has a different log format
	count := len(res) / slowLogEntrySize

	return count, nil
}


// GetCPUUsage retrieves Redis CPU usage statistics.
func (r *RedisMonitor) GetCPUUsage() (float64, error) {
	cpuInfo, err := r.Client.Info(ctx, "cpu").Result()
	if err != nil {
		return 0, err
	}

	lines := strings.Split(cpuInfo, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "used_cpu_sys:") {
			var cpuUsage float64
			_, err := fmt.Sscanf(line, "used_cpu_sys:%f", &cpuUsage)
			if err != nil {
				log.Println("Error parsing CPU usage:", err)
				return 0, err
			}
			return cpuUsage, nil
		}
	}
	return 0, fmt.Errorf("CPU usage metric not found")
}
