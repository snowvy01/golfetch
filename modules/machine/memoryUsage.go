package machine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetMemoryUsage() string {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return "unknown"
	}
	defer file.Close()

	var total, available uint64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "MemTotal:") {
			total = parseMemLine(line)
		} else if strings.HasPrefix(line, "MemAvailable:") {
			available = parseMemLine(line)
		}
		if total > 0 && available > 0 {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "unknown"
	}

	if total == 0 {
		return "unknown"
	}
	used := total - available

	totalGiB := float64(total) / 1024 / 1024
	usedGiB := float64(used) / 1024 / 1024
	percent := (float64(used) / float64(total)) * 100

	return fmt.Sprintf("%.2f GiB / %.2f GiB (%.0f%%)", usedGiB, totalGiB, percent)
}

func parseMemLine(line string) uint64 {
	fields := strings.Fields(line)
	if len(fields) < 2 {
		return 0
	}
	val, _ := strconv.ParseUint(fields[1], 10, 64)
	return val
}
