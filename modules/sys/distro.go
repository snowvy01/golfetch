package sys

import (
	"bufio"
	"os"
	"strings"
)

type info struct {
	NAME     string
	PRNAME   string
	BUILD_ID string
	LOGO     string
}

func GetDistribution() info {
	infor := info{NAME: "Unknown OS", PRNAME: "Unknown", BUILD_ID: "", LOGO: "base"}
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
			infor.NAME = value
		case "PRETTY_NAME":
			infor.PRNAME = value
		case "BUILD_ID":
			infor.BUILD_ID = value
		case "LOGO":
			infor.LOGO = value
		}
	}

	if err := scanner.Err(); err != nil {
		return infor
	}
	return infor
}
