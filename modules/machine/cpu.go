package machine

import (
	"bytes"
	"syscall"
)

func GetCPU() string {
	var buf [2048]byte

	fd, err := syscall.Open("/proc/cpuinfo", syscall.O_RDONLY, 0)
	if err != nil {
		return "Unknown"
	}
	defer syscall.Close(fd)

	n, err := syscall.Read(fd, buf[:])
	if err != nil || n == 0 {
		return "Unknown"
	}

	data := buf[:n]
	prefix := []byte("model name")

	idx := bytes.Index(data, prefix)
	if idx == -1 {
		return "Unknown"
	}

	lineEnd := bytes.IndexByte(data[idx:], '\n')
	if lineEnd == -1 {
		lineEnd = len(data) - idx
	}
	line := data[idx : idx+lineEnd]

	colonIdx := bytes.IndexByte(line, ':')
	if colonIdx == -1 || colonIdx+1 >= len(line) {
		return "Unknown"
	}

	val := line[colonIdx+1:]
	for len(val) > 0 && (val[0] == ' ' || val[0] == '\t') {
		val = val[1:]
	}

	w := 0
	for i := 0; i < len(val); {
		if i+3 <= len(val) && bytes.Equal(val[i:i+3], []byte("(R)")) {
			i += 3
			continue
		}
		if i+4 <= len(val) && bytes.Equal(val[i:i+4], []byte("(TM)")) {
			i += 4
			continue
		}
		val[w] = val[i]
		w++
		i++
	}
	finalBytes := val[:w]

	for len(finalBytes) > 0 && (finalBytes[len(finalBytes)-1] == ' ' || finalBytes[len(finalBytes)-1] == '\r') {
		finalBytes = finalBytes[:len(finalBytes)-1]
	}

	if len(finalBytes) == 0 {
		return "Unknown"
	}

	return string(finalBytes)
}
