package machine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	// Skip the header line ("Filename Type Size Used Priority")
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
