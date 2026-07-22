package machine

import (
	"bufio"
	"os"
	"strings"
)

type FSInfo struct {
	Device     string
	MountPoint string
	FSType     string
}

func GetFSSystems() []FSInfo {
	file, err := os.Open("/proc/mounts")
	if err != nil {
		return nil
	}
	defer file.Close()

	var mounts []FSInfo
	seenDevices := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}

		device := fields[0]
		mountedTo := fields[1]
		fstype := fields[2]

		if !strings.HasPrefix(device, "/dev/") {
			continue
		}
		if fstype == "devtmpfs" {
			continue
		}
		if seenDevices[device] {
			continue
		}

		seenDevices[device] = true
		mounts = append(mounts, FSInfo{
			Device:     device,
			MountPoint: mountedTo,
			FSType:     fstype,
		})
	}
	if scanner.Err() != nil {
		return nil
	}
	return mounts
}
