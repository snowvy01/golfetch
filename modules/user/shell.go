package user

import (
	"os"
	"os/exec"
	"strings"
)

func GetShell() []string {
	// Shell name part:
	var shell []string
	shellPath := os.Getenv("SHELL")
	if shellPath == "" {
		return append(shell, "unknown", "")
	}
	parts := strings.Split(shellPath, "/")
	shellName := parts[len(parts)-1]

	// Shell version part:
	cmd := exec.Command(shellName, "--version")
	output, err := cmd.Output()
	if err != nil {
		return append(shell, shellName, "")
	}
	outStr := string(output)
	fields := strings.Fields(outStr)

	for _, field := range fields {
		if len(field) > 0 && field[0] >= '0' && field[0] <= '9' {
			if idx := strings.Index(field, "("); idx != -1 {
				return append(shell, shellName, field[:idx])
			}
			return append(shell, shellName, field)
		}
	}
	return append(shell, shellName, "")
}
