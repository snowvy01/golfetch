package user

import (
	"os"
)

func GetLocale() string {
	locale := os.Getenv("LANG")
	if locale == "" {
		return "unknown"
	}
	return locale
}
