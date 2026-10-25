package user

import (
	"os"
	"strings"
)

// terminals maps an environment variable that a terminal sets to its name.
var terminals = []struct{ env, name string }{
	{"KITTY_WINDOW_ID", "kitty"},
	{"ALACRITTY_SOCKET", "alacritty"},
	{"ALACRITTY_LOG", "alacritty"},
	{"WEZTERM_EXECUTABLE", "wezterm"},
	{"GHOSTTY_RESOURCES_DIR", "ghostty"},
	{"KONSOLE_VERSION", "konsole"},
	{"FOOT_SERVER_SOCK", "foot"},
}

func GetTerminal() string {
	name := os.Getenv("TERM_PROGRAM")
	version := os.Getenv("TERM_PROGRAM_VERSION")

	if name == "" {
		for _, t := range terminals {
			if os.Getenv(t.env) != "" {
				name = t.name
				break
			}
		}
	}
	if name == "" {
		// TERM is like "xterm-kitty", "foot", "alacritty"
		term := os.Getenv("TERM")
		switch {
		case strings.HasPrefix(term, "xterm-"):
			name = strings.TrimPrefix(term, "xterm-")
		case term != "" && term != "linux" && term != "dumb":
			name = term
		}
	}
	if name == "" {
		return "Unknown"
	}
	name = strings.TrimSuffix(strings.ToLower(name), ".app")
	if name == "apple_terminal" {
		name = "terminal"
	}

	if version == "" {
		version = termVersion(name)
	}
	if version == "" {
		return name
	}
	return name + " " + version
}

func termVersion(name string) string {
	if v := os.Getenv("KONSOLE_VERSION"); v != "" && name == "konsole" && len(v) == 6 {
		// KONSOLE_VERSION is MMmmpp, e.g. 240802 -> 24.8.2
		return trimZero(v[0:2]) + "." + trimZero(v[2:4]) + "." + trimZero(v[4:6])
	}

	files, err := os.ReadDir("/var/lib/pacman/local")
	if err != nil {
		return ""
	}
	prefix := name + "-"
	for _, file := range files {
		n := file.Name()
		if !file.IsDir() || !strings.HasPrefix(n, prefix) {
			continue
		}
		ver := n[len(prefix):]
		// Skip packages like kitty-terminfo-0.49.2-1: version must start with a digit
		if ver == "" || ver[0] < '0' || ver[0] > '9' {
			continue
		}
		if lastDash := strings.LastIndex(ver, "-"); lastDash != -1 {
			ver = ver[:lastDash]
		}
		return ver
	}
	return ""
}

func trimZero(s string) string {
	if t := strings.TrimLeft(s, "0"); t != "" {
		return t
	}
	return "0"
}
