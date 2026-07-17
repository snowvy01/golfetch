package user

import "os"

func GetHome() string {
	home, err := os.UserHomeDir()
	if err == nil {
		return home
	} else {
		envHome := os.Getenv("HOME")
		if envHome != "" {
			return envHome
		} else {
			return "unknown"
		}
	}
}

func GetPWD() string {
	pwd, err := os.Getwd()
	if err != nil {
		return "unknown"
	}
	return pwd
}
