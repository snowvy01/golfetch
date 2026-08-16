package machine

import (
	"bytes"
	"fmt"
	"syscall"
)

func GetMemoryUsage() string {
	var buf [2048]byte

	fd, err := syscall.Open("/proc/meminfo", syscall.O_RDONLY, 0)
	if err != nil {
		return "unknown"
	}
	defer syscall.Close(fd)

	n, err := syscall.Read(fd, buf[:])
	if err != nil || n == 0 {
		return "unknown"
	}

	data := buf[:n]

	totalKB := parseMemBytes(data, []byte("MemTotal:"))
	availableKB := parseMemBytes(data, []byte("MemAvailable:"))

	if totalKB == 0 || availableKB == 0 || totalKB < availableKB {
		return "Unknown"
	}

	usedKB := totalKB - availableKB

	totalGiB := float64(totalKB) / 1024 / 1024
	usedGiB := float64(usedKB) / 1024 / 1024
	percent := (float64(usedKB) / float64(totalKB)) * 100

	return sprintfMem(usedGiB, totalGiB, percent)
}

func parseMemBytes(data []byte, prefix []byte) uint64 {
	idx := bytes.Index(data, prefix)
	if idx == -1 {
		return 0
	}

	ptr := data[idx+len(prefix):]

	i := 0
	for i < len(ptr) && (ptr[i] == ' ' || ptr[i] == '\t') {
		i++
	}

	var val uint64 = 0
	for i < len(ptr) && ptr[i] >= '0' && ptr[i] <= '9' {
		val = val*10 + uint64(ptr[i]-'0')
		i++
	}

	return val
}

func sprintfMem(used, total, percent float64) string {
	return appendFormat(used, total, percent)
}

func appendFormat(used, total, percent float64) string {
	return fmt.Sprintf("%.2f GiB / %.2f GiB \033[35m(%.0f%%)\033[0m", used, total, percent)
}
