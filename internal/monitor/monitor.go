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

// NewRedisMonitor initializes a new Redis client with configurations.
func NewRedisMonitor(redisAddr string) *RedisMonitor {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	return &RedisMonitor{Client: client}
}

// GetMemoryUsage retrieves Redis memory usage in bytes.
func (r *RedisMonitor) GetMemoryUsage() (int64, error) {
	mem, err := r.Client.Info(ctx, "memory").Result()
	if err != nil {
		return 0, err
	}

	lines := strings.Split(mem, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "used_memory:") {
			var usage int64
			_, err := fmt.Sscanf(line, "used_memory:%d", &usage)
			if err != nil {
				log.Println("Error parsing memory usage:", err)
				return 0, err
			}
			return usage, nil
		}
	}
	return 0, fmt.Errorf("memory usage metric not found")
}

// GetSlowQueryCount retrieves the number of slow queries logged by Redis.
func (r *RedisMonitor) GetSlowQueryCount(ctx context.Context) (int, error) {
	res, err := r.Client.Do(ctx, "SLOWLOG", "LEN").Int()
	if err != nil {
		return 0, err
	}
	return res, nil
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
