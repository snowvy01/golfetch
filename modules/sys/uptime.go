package sys

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetUptime() string {
	info, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return "unknown"
	}
	parts := strings.Fields(string(info))
	if len(parts) == 0 {
		return "unknown"
	}

	seconds := parts[0]
	if dotIdx := strings.Index(seconds, "."); dotIdx != -1 {
		seconds = seconds[:dotIdx]
	}

	totalSeconds, err := strconv.Atoi(seconds)
	if err != nil {
		return "unknown"
	}
	hs := totalSeconds / 3600
	mins := (totalSeconds % 3600) / 60
	if hs > 0 {
		return fmt.Sprintf("%d Hours, %d Minutes", hs, mins)
	}
	return fmt.Sprintf("%d Minutes", mins)
}
