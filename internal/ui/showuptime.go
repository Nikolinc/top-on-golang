package ui

import (
	"fmt"
	"top-on-golang/internal/system"
	"time"
)

func ShowUptime(loadavg system.LoadavgStat, uptime system.UptimeStat) {

	fmt.Printf("top - %8s up %d days, %2d:%2d 2 users, load average: %.2f %.2f %.2f\n",
		time.Now().Format("15:04:05"), uptime.Days, uptime.Hours, uptime.Minutes, loadavg.LoadavgOne, loadavg.LoadavgFive, loadavg.LoadavgFifteen)

}
