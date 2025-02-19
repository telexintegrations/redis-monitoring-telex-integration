package monitor

import (
	"fmt"
)

const (
	MemoryThreshold  = 100 * 1024 * 1024 // 100MB
	SlowLogThreshold = 10                // Alert if slow queries exceed this
	CPUThreshold     = 80.0              // Alert if CPU usage exceeds 80%
)

func ShouldSendAlert(memUsage int64, slowLogs int, cpuUsage float64) string {
	if memUsage > MemoryThreshold {
		return fmt.Sprintf("🚨 High Memory Usage: %s", formatBytes(memUsage))
	} else if slowLogs > SlowLogThreshold {
		return fmt.Sprintf("⚠️ High Slow Query Count: %d", slowLogs)
	} else if cpuUsage > CPUThreshold {
		return fmt.Sprintf("🔥 High CPU Usage: %.2f%%", cpuUsage)
	}
	return ""
}

func formatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}
