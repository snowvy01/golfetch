// Copyright (c) 2026, Snowvy. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or project headers.

package main

import (
	// These are basic packages. You can remove them if you do not need them (Be sure if program will work without them before deleting):
	"fmt"
	//"log"
	//"os"
	"strings"

	// Packages from golfetch/modules (Change them how do you want):
	"golfetch/modules/sys"
	//"golfetch/modules/logo"
	"golfetch/modules/machine"
	"golfetch/modules/user"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
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

	// Currently, there's no functions in logo package :P

	// Main output:
	usrhost := fmt.Sprintf("%s%s%s@%s%s%s", Cyan, userinfo[0], Reset, Cyan, hostname, Reset) // user and hostname: username@hostname
	fmt.Println(usrhost)
	fmt.Println(strings.Repeat("-", len(userinfo[0]+hostname)+1))                          // spacer
	fmt.Printf("~ %sOS%s: %s\n", Cyan, Reset, osinfo.NAME)                                 // OS name (from /etc/os-release)
	fmt.Printf("~ %sHost%s: %s\n", Cyan, Reset, realhost)                                  // Real hostname (from /sys/class/dmi/id/product_name)
	fmt.Printf("~ %sKernel%s: %s\n", Cyan, Reset, kernel)                                  // Linux kernel version
	fmt.Printf("~ %sUptime%s: %s\n", Cyan, Reset, uptime)                                  // Uptime (00 hours, 00 minutes)
	fmt.Printf("~ %sPackages%s: %s\n", Cyan, Reset, pkgs)                                  // Count of all installed packages (pacman)
	fmt.Printf("~ %sShell%s: %s\n", Cyan, Reset, fmt.Sprintf("%s %s", shell[0], shell[1])) // Current terminal shell (Bash, zsh, fish) and it's version
	fmt.Printf("~ %sCPU%s: %s\n", Cyan, Reset, cpuName)                                    // Model name of CPU
	fmt.Printf("~ %sMemory%s: %s\n", Cyan, Reset, memUsage)                                // Memory usage
	fmt.Printf("~ %sSwap%s: %.2f GiB / %.2f (%s)\n", Cyan, Reset, swapInfo.Used, swapInfo.Size, swapInfo.Name)
	fmt.Printf("~ %sHome%s: %s\n", Cyan, Reset, home)     // Home directory of current user
	fmt.Printf("~ %sPWD%s: %s\n", Cyan, Reset, pwd)       // Current directory (PWD)
	fmt.Printf("~ %sLocale%s: %s\n", Cyan, Reset, locale) // Locale (current language, by default it's C.UTF-8 or en_US.UTF-8)
	for _, mount := range mounts {
		fmt.Printf("~ %sDrive (%s - %s)%s: %s\n", Cyan, mount.Device, mount.MountPoint, Reset, mount.FSType)
	}
}
