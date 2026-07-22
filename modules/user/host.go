package user

import (
	"fmt"
	"os"
	"strings"
)

func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hostname
}

func GetRealHostname() string {
	mainName, err := os.ReadFile("/sys/class/dmi/id/product_name")
	normalmainName := strings.TrimSpace(string(mainName))
	if err != nil || normalmainName == "" {
		return "unknown"
	}
	verName, err := os.ReadFile("/sys/class/dmi/id/product_version")
	normalverName := strings.TrimSpace(string(verName))
	if err != nil || normalverName == "" {
		return normalmainName
	}
	return fmt.Sprintf("%s (%s)", normalmainName, normalverName)
}
