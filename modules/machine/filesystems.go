package machine

import (
	"bufio"
	"os"
	"regexp"
	"strings"
	"syscall"
)

type FSInfo struct {
	Device     string
	MountPoint string
	FSType     string
	Used       float64
	Total      float64
	Percentage float64
}

func GetFSSystems() []FSInfo {
	file, err := os.Open("/proc/mounts")
	if err != nil {
		return []FSInfo{}
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

		re := regexp.MustCompile(`^(/efi|/apex|/bootstrap-apex|/cache|/efs|/product)`)

		if re.MatchString(mountedTo) {
			continue
		}
		if fstype == "devtmpfs" {
			continue
		}
		if seenDevices[device] {
			continue
		}

		var stat syscall.Statfs_t
		if err := syscall.Statfs(mountedTo, &stat); err != nil {
			continue
		}

		const gib = 1024 * 1024 * 1024
		totalBytes := stat.Blocks * uint64(stat.Bsize)
		freeBytes := stat.Bavail * uint64(stat.Bsize)
		usedBytes := totalBytes - freeBytes

		var percent float64
		if totalBytes > 0 {
			percent = (float64(usedBytes) / float64(totalBytes)) * 100
		}

		seenDevices[device] = true
		mounts = append(mounts, FSInfo{
			Device:     device,
			MountPoint: mountedTo,
			FSType:     fstype,
			Used:       float64(usedBytes) / gib,
			Total:      float64(totalBytes) / gib,
			Percentage: percent,
		})
	}
	if scanner.Err() != nil {
		return []FSInfo{}
	}
	return mounts
}
