package logo

import (
	"fmt"
	"math"
	"strings"
)

func MakeItRainbow(text string) string {
	var builder strings.Builder
	builder.Grow(len(text) * 20)

	frequency := 0.1
	for i, r := range text {
		if r == '\n' || r == '\r' {
			builder.WriteRune(r)
			continue
		}

		red := int(math.Sin(frequency*float64(i)+0)*127 + 128)
		green := int(math.Sin(frequency*float64(i)+2)*127 + 128)
		blue := int(math.Sin(frequency*float64(i)+4)*127 + 128)

		fmt.Fprintf(&builder, "\x1b[38;2;%d;%d;%dm%c\x1b[0m", red, green, blue, r)
	}

	return builder.String()
}
