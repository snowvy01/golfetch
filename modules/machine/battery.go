package machine

import (
	"os"
	"path/filepath"
	"strings"
)

func GetBattery() [][]string {
	var batteries [][]string
	baseDir := "/sys/class/power_supply"

	entries, err := os.ReadDir(baseDir)
	if err != nil {
		return batteries
	}

	for _, entry := range entries {
		dirPath := filepath.Join(baseDir, entry.Name())

		typeBytes, err := os.ReadFile(filepath.Join(dirPath, "type"))
		if err != nil {
			continue
		}

		if strings.TrimSpace(string(typeBytes)) != "Battery" {
			continue
		}

		// Собираем сырые данные без форматирования
		batteryData := []string{
			readSysFile(dirPath, "capacity", "0"),
			readSysFile(dirPath, "status", "Unknown"),
			readSysFile(dirPath, "model_name", "Generic"),
			readSysFile(dirPath, "manufacturer", "Unknown"),
			readSysFile(dirPath, "technology", "Li-ion"),
		}

		batteries = append(batteries, batteryData)
	}

	return batteries
}

func readSysFile(dir, file, fallback string) string {
	bytes, err := os.ReadFile(filepath.Join(dir, file))
	if err != nil {
		return fallback
	}
	return strings.TrimSpace(string(bytes))
}
