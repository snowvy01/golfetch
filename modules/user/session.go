package user

import (
	"os"
)

func GetSession() (Name string, Type string) {
	name := os.Getenv("XDG_CURRENT_DESKTOP")
	stype := os.Getenv("XDG_SESSION_TYPE")
	if name != "" && stype != "" {
		return name + " -", stype
	}
	return "unknown", ""
}
