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

// Ansi escapes
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
	// Text styles
	Bold   = "\033[1m" // Bold text. To reset it, use Reset string
	Italic = "\x1b[3m" // Italic text.
)

func main() {
	osinfo := sys.GetDistribution()
	kernel := sys.GetKernelVer()
	uptime := sys.GetUptime()
	pkgs := sys.GetPacmanPkgs()

	cpuName := machine.GetCPU()
	memUsage := machine.GetMemoryUsage()
	swapsInfo := machine.GetSwap()

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

	usrhost := fmt.Sprintf("%s%s%s@%s%s%s", Italic+bCyan, userinfo[0], Cyan, Italic+bCyan, hostname, Reset) // user and hostname spaced by @
	spacer := strings.Repeat("~", len(userinfo[0]+hostname)+1)
	infoLines = append(
		infoLines,
		usrhost,
		spacer,
		fmt.Sprintf("> %sOS:%s %s - %s %s(%s)%s", Bold+Cyan, Reset, osinfo.PrName, osinfo.BuildID, Magenta, osinfo.Arch, Reset),
		fmt.Sprintf("> %sHost:%s %s at %s", Bold+Cyan, Reset, userinfo[1], realhost),
		fmt.Sprintf("> %sKernel:%s %s", Bold+Cyan, Reset, kernel),
		fmt.Sprintf("> %sUptime:%s %s", Bold+Cyan, Reset, uptime),
		fmt.Sprintf("> %sPackages:%s %s", Bold+Cyan, Reset, pkgs),
		fmt.Sprintf("> %sSession:%s %s %s", Bold+Cyan, Reset, sesName, sesType),
		fmt.Sprintf("> %sShell:%s %s", Bold+Cyan, Reset, shell),
		fmt.Sprintf("> %sCPU:%s %s", Bold+Cyan, Reset, cpuName),
		fmt.Sprintf("> %sMemory:%s %s", Bold+Cyan, Reset, memUsage))
	if len(swapsInfo) == 0 {
		infoLines = append(infoLines, fmt.Sprintf("> %sSwap:%s %sDisabled%s", Bold+Cyan, Reset, bRed, Reset))
	} else {
		for _, swap := range swapsInfo {
			infoLines = append(infoLines, fmt.Sprintf("> %sSwap %s(%s)%s:%s %.2f GiB / %.2f GiB %s(%s)%s",
				Bold+Cyan, Magenta, swap.Name, Reset+Bold+Cyan, Reset, swap.Used, swap.Size, Magenta, swap.Percent, Reset))
		}
	}
	infoLines = append(
		infoLines,
		fmt.Sprintf("> %sHome:%s %s", Bold+Cyan, Reset, home),
		fmt.Sprintf("> %sPWD:%s %s", Bold+Cyan, Reset, pwd),
		fmt.Sprintf("> %sLocale:%s %s", Bold+Cyan, Reset, locale),
	)
	for _, mount := range mounts {
		infoLines = append(infoLines, fmt.Sprintf("> %sDrive %s(%s, %s)%s:%s %.2f GiB / %.2f GiB %s(%.0f%%)%s",
			Bold+Cyan, Magenta, mount.MountPoint, mount.FSType, Reset+Cyan+Bold, Reset, mount.Used, mount.Total, Magenta, mount.Percentage, Reset))
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
