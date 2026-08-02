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

// Ansi escapes colors
const (
	// Normal colors
	Reset   = "\033[0m"
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	// Brighter colors
	bBlack   = "\x1b[90m"
	bRed     = "\x1b[91m"
	bGreen   = "\x1b[92m"
	bYellow  = "\x1b[93m"
	bBlue    = "\x1b[94m"
	bMagenta = "\x1b[95m"
	bCyan    = "\x1b[96m"
	bWhite   = "\x1b[97m"
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
	sesName, sesType := user.GetSession()
	shell := user.GetShell()
	home := user.GetHome()
	pwd := user.GetPWD()
	locale := user.GetLocale()
	mounts := machine.GetFSSystems()

	var infoLines []string

	usrhost := fmt.Sprintf("%s%s%s@%s%s%s", bCyan, userinfo[0], Cyan, bCyan, hostname, Reset) // user and hostname spaced by @
	spacer := strings.Repeat("-", len(userinfo[0]+hostname)+1)
	infoLines = append(
		infoLines,
		usrhost,
		spacer,
		fmt.Sprintf("~ %sOS%s: %s - %s", Cyan, Reset, osinfo.Name, osinfo.BuildID),
		fmt.Sprintf("~ %sHost%s: %s at %s", Cyan, Reset, userinfo[1], realhost),
		fmt.Sprintf("~ %sKernel%s: %s", Cyan, Reset, kernel),
		fmt.Sprintf("~ %sUptime%s: %s", Cyan, Reset, uptime),
		fmt.Sprintf("~ %sPackages%s: %s", Cyan, Reset, pkgs),
		fmt.Sprintf("~ %sSession%s: %s %s", Cyan, Reset, sesName, sesType),
		fmt.Sprintf("~ %sShell%s: %s %s", Cyan, Reset, shell[0], shell[1]),
		fmt.Sprintf("~ %sCPU%s: %s", Cyan, Reset, cpuName),
		fmt.Sprintf("~ %sMemory%s: %s", Cyan, Reset, memUsage),
		fmt.Sprintf("~ %sSwap (%s)%s: %.2f GiB / %.2f GiB (%s)", Cyan, swapInfo.Name, Reset, swapInfo.Used, swapInfo.Size, swapInfo.Percent),
		fmt.Sprintf("~ %sHome%s: %s", Cyan, Reset, home),
		fmt.Sprintf("~ %sPWD%s: %s", Cyan, Reset, pwd),
		fmt.Sprintf("~ %sLocale%s: %s", Cyan, Reset, locale),
	)
	if mounts != nil {
		for _, mount := range mounts {
			infoLines = append(infoLines, fmt.Sprintf("~ %sDrive (%s, %s)%s: %.2f GiB / %.2f GiB (%.0f%%)",
				Cyan, mount.MountPoint, mount.FSType, Reset, mount.Used, mount.Total, mount.Percentage))
		}
	}
	infoLines = append(infoLines, "",
		fmt.Sprintf("%s███%s███%s███%s███%s███%s███%s███%s███%s", Black, Red, Green, Yellow, Blue, Magenta, Cyan, White, Reset),
		fmt.Sprintf("%s███%s███%s███%s███%s███%s███%s███%s███%s", bBlack, bRed, bGreen, bYellow, bBlue, bMagenta, bCyan, bWhite, Reset))

	logoLines := logo.GetLogo()
	maxLogoWidth := 0
	for _, line := range logoLines {
		if len(line) > maxLogoWidth {
			maxLogoWidth = len(line)
		}
	}

	padding := "  "

	maxLines := max(len(logoLines), len(infoLines))

	for i := range maxLines {
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
