package user

import (
	"os"
	"os/user"
)

func GetUserInfo() []string {
	currentUser, err := user.Current()
	if err == nil {
		if currentUser.Name != "" {
			return []string{currentUser.Username, currentUser.Name}
		} else {
			return []string{currentUser.Username, currentUser.Username}
		}
	} else {
		if envStr := os.Getenv("USER"); envStr != "" {
			return []string{envStr, envStr}
		}
	}
	return []string{"unknown", "Unknown"}
}
