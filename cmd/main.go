package main

import (
	"top-on-golang/internal/system"
	"top-on-golang/internal/ui"
)

func main() {
	var meminfo system.MemStat = system.MemInfo()
	var loadavg system.LoadavgStat = system.Loadavg()
	var uptime system.UptimeStat = system.Uptime()
	var cpu system.CPUState = system.CPU()

	ui.ShowUptime(loadavg, uptime)
	ui.ShowCPU(cpu)
	ui.ShowMem(meminfo)
}