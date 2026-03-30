package system

import (
	"io"
	"os"
	"strings"
	"strconv"
)


type MemStat struct {
	Total, Used, Free, BuffCache, Available, SwapTotal, SwapFree, SwapUsed float64
}


func GetVolume(data []byte, count int) float64{
	line := strings.Split(string(data), "\n")[count]
	fields := strings.Fields(line)[1:]
	if len(fields) < 2 {
		return 0
	}
	val,_ := strconv.ParseFloat(fields[0],64)
	return val
}

func MemInfo() MemStat {
	file, _ := os.Open("/proc/meminfo")
	defer file.Close()
	data, _ := io.ReadAll(file)

	var total, used, free, buffcache, available float64
	var swaptotal, swapfree, swapused float64


	total = GetVolume(data, 0)/1024
	free = GetVolume(data, 1)/1024
	available = GetVolume(data, 2)/1024
	used = total - available
	buffcache = GetVolume(data, 3)/1024 + GetVolume(data, 4)/1024

	swaptotal = GetVolume(data, 14)/1024
	swapfree = GetVolume(data, 15)/1024
	swapused =  swaptotal - swapfree

	return MemStat{
		Total: total, Used: used, Free: free, BuffCache: buffcache,
		Available: available, SwapTotal: swaptotal, SwapFree: swapfree, SwapUsed: swapused,
	}
}