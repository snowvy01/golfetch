package machine

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func GetMemoryUsage() string {
	var buf [2048]byte

	fd, err := syscall.Open("/proc/meminfo", syscall.O_RDONLY, 0)
	if err != nil {
		return "Unknown"
	}
	defer syscall.Close(fd)

	n, err := syscall.Read(fd, buf[:])
	if err != nil || n == 0 {
		return "Unknown"
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

type Swapinfo struct {
	Name    string
	Size    float64
	Used    float64
	Percent string
}

func GetSwap() []Swapinfo {
	file, err := os.Open("/proc/swaps")
	if err != nil {
		return []Swapinfo{}
	}
	defer file.Close()

	var swaps []Swapinfo
	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return []Swapinfo{}
	}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 4 {
			continue
		}

		sizeKB, err1 := strconv.ParseFloat(fields[2], 64)
		usedKB, err2 := strconv.ParseFloat(fields[3], 64)
		if err1 != nil || err2 != nil {
			continue
		}

		var percent string
		if sizeKB > 0 {
			percent = fmt.Sprintf("%.0f%%", (usedKB/sizeKB)*100)
		} else {
			percent = "0%"
		}

		swaps = append(swaps, Swapinfo{
			Name:    fields[0],
			Size:    sizeKB / (1024 * 1024), // Converts KB to GB
			Used:    usedKB / (1024 * 1024), // Converts KB to GB
			Percent: percent,
		})
	}

	if scanner.Err() != nil {
		return []Swapinfo{}
	}

	return swaps
}

func sprintfMem(used, total, percent float64) string {
	return appendFormat(used, total, percent)
}

func appendFormat(used, total, percent float64) string {
	return fmt.Sprintf("%.2f GiB / %.2f GiB \033[35m(%.0f%%)\033[0m", used, total, percent)
}
