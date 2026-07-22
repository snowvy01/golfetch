package sys

import (
	"os"
	"strconv"
)

func GetPacmanPkgs() string {
	files, err := os.ReadDir("/var/lib/pacman/local")
	if err != nil {
		return "0 (pacman)"
	}
	var count int
	for _, file := range files {
		if file.IsDir() {
			count++
		}
	}
	return strconv.Itoa(count) + " (pacman)"
}
