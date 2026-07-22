// Copyright (c) 2026, Snowvy. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or project headers.

package main

import (
	// These are basic packages. You can remove them if you do not need them (Be sure if program will work without them before deleting):
	"fmt"
	"strings"

	// Packages from golfetch/modules (Change them how do you want):
	"golfetch/modules/logo"
	"golfetch/modules/machine"
	"golfetch/modules/sys"
	"golfetch/modules/user"
)

const (
	Reset = "\033[0m"
	//Red     = "\033[31m"
	//Green = "\033[32m"
	//Yellow  = "\033[33m"
	//Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
)

func main() {
	osinfo := sys.GetDistribution()
	kernel := sys.GetKernelVer()
	uptime := sys.GetUptime()
	pkgs := sys.GetPacmanPkgs()

	cpuName := machine.GetCPU()
	memUsage := machine.GetMemoryUsage()
	swapInfo := machine.GetSwap()

	userinfo := user.GetUserInfo()
	hostname := user.GetHostname()
	realhost := user.GetRealHostname()
	shell := user.GetShell()
	home := user.GetHome()
	pwd := user.GetPWD()
	locale := user.GetLocale()
	mounts := machine.GetFSSystems()

	var infoLines []string

	usrhost := fmt.Sprintf("%s%s%s@%s%s%s", Magenta, userinfo[0], Reset, Magenta, hostname, Reset) // user and hostname spaced by @
	spacer := strings.Repeat("-", len(userinfo[0]+hostname)+1)
	infoLines = append(
		infoLines,
		usrhost,
		spacer,
		fmt.Sprintf("~ %sOS%s: %s (%s)", Cyan, Reset, osinfo.Name, osinfo.BuildID),
		fmt.Sprintf("~ %sHost%s: %s", Cyan, Reset, realhost),
		fmt.Sprintf("~ %sKernel%s: %s", Cyan, Reset, kernel),
		fmt.Sprintf("~ %sUptime%s: %s", Cyan, Reset, uptime),
		fmt.Sprintf("~ %sPackages%s: %s", Cyan, Reset, pkgs),
		fmt.Sprintf("~ %sShell%s: %s %s", Cyan, Reset, shell[0], shell[1]),
		fmt.Sprintf("~ %sCPU%s: %s", Cyan, Reset, cpuName),
		fmt.Sprintf("~ %sMemory%s: %s", Cyan, Reset, memUsage),
		fmt.Sprintf("~ %sSwap%s: %.2f GiB / %.2f (%s)", Cyan, Reset, swapInfo.Used, swapInfo.Size, swapInfo.Name),
		fmt.Sprintf("~ %sHome%s: %s", Cyan, Reset, home),
		fmt.Sprintf("~ %sPWD%s: %s", Cyan, Reset, pwd),
		fmt.Sprintf("~ %sLocale%s: %s", Cyan, Reset, locale),
	)
	for _, mount := range mounts {
		infoLines = append(infoLines, fmt.Sprintf("~ %sDrive (%s - %s)%s: %s", Cyan, mount.Device, mount.MountPoint, Reset, mount.FSType))
	}

	logoLines := logo.GetLogo()
	maxLogoWidth := 0
	for _, line := range logoLines {
		if len(line) > maxLogoWidth {
			maxLogoWidth = len(line)
		}
	}

	padding := "  "

	maxLines := len(logoLines)
	if len(infoLines) > maxLines {
		maxLines = len(infoLines)
	}
	for i := 0; i < maxLines; i++ {
		logoPart := ""
		infoPart := ""

		if i < len(logoLines) {
			logoPart = Cyan + logoLines[i] + Reset
			spaceCount := maxLogoWidth - len(logoLines[i])
			logoPart += strings.Repeat(" ", spaceCount)
		} else {
			logoPart = strings.Repeat(" ", maxLogoWidth)
		}

		if i < len(infoLines) {
			infoPart = infoLines[i]
		}

		fmt.Printf("%s%s%s\n", logoPart, padding, infoPart)
	}
}
