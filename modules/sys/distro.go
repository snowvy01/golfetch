package sys

import (
	"bufio"
	"os"
	"strings"
)

type info struct {
	Name    string
	PrName  string
	BuildID string
	Arch    string
}

func GetDistribution() info {
	infor := info{Name: "Unknown OS", PrName: "Unknown", BuildID: "", Arch: ""}
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return infor
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			if err != nil {
				err = closeErr
			}
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := parts[0]
		value := strings.Trim(parts[1], `"'`)

		switch key {
		case "NAME":
			infor.Name = value
		case "PRETTY_NAME":
			infor.PrName = value
		case "BUILD_ID":
			infor.BuildID = value
		}
	}

	if err := scanner.Err(); err != nil {
		return infor
	}

	infor.Arch = getArch()

	return infor
}

func getArch() string {
	content, err := os.ReadFile("/proc/sys/kernel/arch")
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(content))
}
