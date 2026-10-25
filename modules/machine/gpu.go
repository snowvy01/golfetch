package machine

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var pciIDsPaths = []string{
	"/usr/share/hwdata/pci.ids",
	"/usr/share/misc/pci.ids",
	"/usr/share/pci.ids",
}

// GetGPUs returns a name for every display controller found in /sys/class/drm.
func GetGPUs() []string {
	cards, err := filepath.Glob("/sys/class/drm/card[0-9]*/device")
	if err != nil {
		return []string{"Unknown"}
	}

	var gpus []string
	seen := make(map[string]bool)
	for _, dev := range cards {
		class := readSysfs(filepath.Join(dev, "class"))
		if !strings.HasPrefix(class, "0x03") {
			continue
		}
		vendor := strings.TrimPrefix(readSysfs(filepath.Join(dev, "vendor")), "0x")
		device := strings.TrimPrefix(readSysfs(filepath.Join(dev, "device")), "0x")
		if vendor == "" || device == "" || seen[vendor+device] {
			continue
		}
		seen[vendor+device] = true
		gpus = append(gpus, lookupPCI(vendor, device))
	}

	if len(gpus) == 0 {
		return []string{"Unknown"}
	}
	return gpus
}

func readSysfs(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return strings.ToLower(strings.TrimSpace(string(data)))
}

func lookupPCI(vendorID, deviceID string) string {
	var file *os.File
	for _, p := range pciIDsPaths {
		f, err := os.Open(p)
		if err == nil {
			file = f
			break
		}
	}
	if file == nil {
		return fmt.Sprintf("%s:%s", vendorID, deviceID)
	}
	defer file.Close()

	vendorName, deviceName := "", ""
	inVendor := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line[0] == '#' {
			continue
		}
		if line[0] != '\t' {
			if inVendor {
				break
			}
			if strings.HasPrefix(line, vendorID+"  ") {
				vendorName = strings.TrimSpace(line[len(vendorID):])
				inVendor = true
			}
			continue
		}
		if inVendor && strings.HasPrefix(line, "\t"+deviceID+"  ") {
			deviceName = strings.TrimSpace(line[len(deviceID)+1:])
			break
		}
	}

	if vendorName == "" {
		return fmt.Sprintf("%s:%s", vendorID, deviceID)
	}
	vendorName = shortVendor(vendorName)
	if deviceName == "" {
		return fmt.Sprintf("%s %s", vendorName, deviceID)
	}
	// pci.ids puts the marketing name in brackets: "GA107M [GeForce RTX 2050]"
	if l, r := strings.Index(deviceName, "["), strings.LastIndex(deviceName, "]"); l != -1 && r > l {
		deviceName = deviceName[l+1 : r]
	}
	return vendorName + " " + deviceName
}

func shortVendor(name string) string {
	switch {
	case strings.Contains(name, "NVIDIA"):
		return "NVIDIA"
	case strings.Contains(name, "Advanced Micro Devices"), strings.Contains(name, "AMD"):
		return "AMD"
	case strings.Contains(name, "Intel"):
		return "Intel"
	}
	for _, suffix := range []string{" Corporation", " Corp.", " Inc.", ", Inc."} {
		name = strings.TrimSuffix(name, suffix)
	}
	return name
}
