package machine

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Swapinfo struct {
	Name string
	Size float64
	Used float64
}

func GetSwap() *Swapinfo {
	file, err := os.Open("/proc/swaps")
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		_ = scanner.Text()
	}
	if !scanner.Scan() {
		return nil
	}
	fields := strings.Fields(scanner.Text())
	if len(fields) < 4 {
		return nil
	}

	sizeKB, _ := strconv.ParseFloat(fields[2], 64)
	usedKB, _ := strconv.ParseFloat(fields[3], 64)

	return &Swapinfo{
		Name: fields[0],
		Size: sizeKB / (1024 * 1024),
		Used: usedKB / (1024 * 1024),
	}
}
