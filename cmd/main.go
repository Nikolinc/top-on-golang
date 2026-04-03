package main

import (
	"time"
	"top-on-golang/internal/system"
	"top-on-golang/internal/ui"
)

func main() {

	var laststate system.CPUState = system.ReadStat()

	for {
		time.Sleep(time.Second)

		var meminfo system.MemStat = system.MemInfo()
		var loadavg system.LoadavgStat = system.Loadavg()
		var uptime system.UptimeStat = system.Uptime()
		var cpu system.CPUState = system.CPU(laststate)
		var tasks system.TasksState
		tasks, _ = system.Processing()

		ui.Clear()

		ui.ShowUptime(loadavg, uptime)
		ui.ShowTasks(tasks)
		ui.ShowCPU(cpu)
		ui.ShowMem(meminfo)
		ui.ShowProceing()
	}
}
