package system

import (
	"os"
	"strings"
	"strconv"
)

type LoadavgStat struct {
	LoadavgOne, LoadavgFive, LoadavgFifteen float64
}

type UptimeStat struct {
	Days, Hours, Minutes int
}


func Loadavg() LoadavgStat {
	data, _ := os.ReadFile("/proc/loadavg")
	fields := strings.Fields(string(data))

	loadavgOne, _ := strconv.ParseFloat(fields[0], 64)
	loadavgFive, _ := strconv.ParseFloat(fields[1], 64)
	loadavgFifteen, _ := strconv.ParseFloat(fields[2], 64)


	return LoadavgStat{ LoadavgOne:loadavgOne, LoadavgFive: loadavgFive, LoadavgFifteen: loadavgFifteen}
}

func Uptime()UptimeStat{

	data, _ := os.ReadFile("/proc/uptime")

	fields := strings.Fields(string(data))
	uptimeSec, _ := strconv.ParseFloat(fields[0], 64)

	var days, hours, minutes int

	days = int(uptimeSec) / 86400
	hours = (int(uptimeSec) % 86400) / 3600
	minutes = (int(uptimeSec) % 3600) / 60

	return UptimeStat{
		Days:days , Hours: hours, Minutes: minutes,
	}
}