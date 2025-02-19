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

// get memory usage 
func (r *RedisMonitor) GetMemoryUsage(ctx context.Context) (string, error) {
	return r.getInfoSection(ctx, "memory")
}

// get CPU usage data
func (r *RedisMonitor) GetCPUUsage(ctx context.Context) (string, error) {
	return r.getInfoSection(ctx, "cpu")
}

// get keyspace size
func (r *RedisMonitor) GetKeyspaceSize(ctx context.Context) (string, error) {
	return r.getInfoSection(ctx, "keyspace")
}

// get command stats
func (r *RedisMonitor) GetCommandStats(ctx context.Context) (string, error) {
	return r.getInfoSection(ctx, "commandstats")
}

// get real time commaandd execution
func (r *RedisMonitor) GetRealTimeCommandExecution(ctx context.Context) ([]string, error) {
	pubsub := r.Client.Subscribe(ctx, "__redis__:monitor")
	defer pubsub.Close()

	messages := []string{}
	for i := 0; i < 5; i++ { // Capture a few real-time commands
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg.Payload)
	}
	return messages, nil
}

// get slow queries
func (r *RedisMonitor) GetSlowQueries(ctx context.Context) ([]string, error) {
	res, err := r.Client.Do(ctx, "SLOWLOG", "GET").StringSlice()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// retrieve specific sections of redis server information
func (r *RedisMonitor) getInfoSection(ctx context.Context, section string) (string, error) {
	info, err := r.Client.Info(ctx, section).Result()
	if err != nil {
		return "", err
	}
	return info, nil
}
