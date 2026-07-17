package machine

import (
	"bufio"
	"os"
	"strings"
)

func GetCPU() string {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return "unknown"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "model name") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				name := strings.TrimSpace(parts[1])
				name = strings.ReplaceAll(name, "(R)", "")
				name = strings.ReplaceAll(name, "(TM)", "")
				return name
			}
		}
	}
	if scanner.Err() != nil {
		return "unknown"
	}
	return "unknown"
}
