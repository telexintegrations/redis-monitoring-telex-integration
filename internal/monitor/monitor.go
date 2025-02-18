package monitor

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisMonitor struct {
	Client *redis.Client
}

func NewRedisMonitor(addr string) *RedisMonitor {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisMonitor{Client: client}
}

func (r *RedisMonitor) GetMemoryUsage(ctx context.Context) (string, error) {
	mem, err := r.Client.Info(ctx, "memory").Result()
	if err != nil {
		return "", err
	}
	return mem, nil
}

func (r *RedisMonitor) GetCPUUsage(ctx context.Context) (string, error) {
	cpu, err := r.Client.Info(ctx, "cpu").Result()
	if err != nil {
		return "", err
	}
	return cpu, nil
}

// Add more metric functions...

// Metrics
	//keyspace size
	// command statistics
	// cpu usage
	// mem usage
	// real time command execution
	// slow queries