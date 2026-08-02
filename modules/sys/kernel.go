package sys

import (
	"os"
	"strings"
)

func GetKernelVer() string {
	krnl, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return "unknown"
	}
	kernel := strings.TrimSpace(string(krnl))
	systype, err := os.ReadFile("/proc/sys/kernel/ostype")
	if err != nil {
		return kernel
	}
	systemtype := strings.TrimSpace(string(systype))
	return systemtype + " " + kernel
}
