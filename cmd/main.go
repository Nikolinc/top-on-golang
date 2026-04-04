package main

import (
	"time"
	"top-on-golang/internal/system"
	"top-on-golang/internal/ui"
	"top-on-golang/internal/system/tasks"
)

func main() {

	var laststate system.CPUState = system.ReadStat()

	ui.Clear()

	for {
		time.Sleep(time.Second)

		var uptime system.UptimeStat = system.Uptime()
		var meminfo system.MemStat = system.MemInfo()
		var loadavg system.LoadavgStat = system.Loadavg()
		var cpu system.CPUState = system.CPU(laststate)
		var tasks tasks.TaskTotalState
		var list []system.ProcessingInitState
		tasks, list = system.Processing()

		ui.ShowUptime(loadavg, uptime, 0)
		ui.ShowTasks(tasks, 1 )
		ui.ShowCPU(cpu, 2)
		ui.ShowMem(meminfo, 3)
		ui.ShowProceing(list,7)

	}
}
