# golfetch

Golfetch is a [neofetch](https://github.com/dylanaraps/neofetch)-like tool for fetching your system information and displaying it in a way you want. It is written in Go, with a focus on performance, customization, and the suckless style. Currently, I maintain it only on Arch-based systems. If you want to use this soft on other distros or systems, change components and compile it

> Note: golfetch is only tested on x86_64 platforms by maintainer (me). It will most likely run on other platforms, but there is no guarantee.

## Installation:

### Manual build
```bash
git clone https://github.com/snowvy01/golfetch.git # clone the repo
cd golfetch
go build -trimpath -buildmode=pie -ldflags="-s -w" -o golfetch # compile it
./golfetch # now execute!
```
### Arch Linux (AUR)
you can install it by using your favorite AUR helper
```bash
yay -S golfetch-git
```
### Configuration
To change the program configuration, modify the main.go file and compile the program:
```bash
go build -trimpath -buildmode=pie -ldflags="-s -w" -o golfetch
```
