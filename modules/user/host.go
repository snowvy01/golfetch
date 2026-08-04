package user

import (
	"os"
	"strings"
)

func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		content, err := os.ReadFile("/etc/hostname")
		if err != nil {
			return "Unknown"
		}
		return strings.TrimSpace(string(content))
	}
	return hostname
}

func GetRealHostname() string {
	content, err := os.ReadFile("/sys/class/dmi/id/product_name")
	if err != nil || string(content) == "" {
		return "unknown"
	}
	return strings.TrimSpace(string(content))
}
