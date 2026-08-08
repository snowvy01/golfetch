package user

import (
	"os"
	"strings"
)

func GetShell() string {
	shellpath := os.Getenv("SHELL")
	if shellpath == "" {
		return "Unknown"
	}

	shellname := shellpath[strings.LastIndex(shellpath, "/")+1:]
	files, err := os.ReadDir("/var/lib/pacman/local")
	if err != nil {
		return shellname
	}

	prefix := shellname + "-"
	for _, file := range files {
		name := file.Name()
		if file.IsDir() && strings.HasPrefix(name, prefix) {
			ver := name[len(prefix):]
			if lastDash := strings.LastIndex(ver, "-"); lastDash != -1 {
				ver = ver[:lastDash]
			}
			return shellname + " " + ver
		}
	}
	return shellname
}
