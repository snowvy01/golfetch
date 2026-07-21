package logo

import (
	_ "embed"
	"strings"
)

//go:embed logo.txt
var logoArch []byte

func GetLogo() []string {
	return strings.Split(string(logoArch), "\n")
}
