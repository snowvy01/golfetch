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
}

func GetDistribution() info {
	infor := info{Name: "Unknown OS", PrName: "Unknown", BuildID: ""}
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return infor
	}
	defer file.Close()

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
	return infor
}
