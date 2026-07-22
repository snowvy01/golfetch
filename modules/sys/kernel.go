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
	return strings.TrimSpace(string(krnl))
}
