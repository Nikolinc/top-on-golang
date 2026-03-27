package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
)

func getvolume(data []byte, count int) float64{
	line := strings.Split(string(data), "\n")[count]
	fields := strings.Fields(line)[1:]
	if len(fields) < 2 {
		return 0
	}
	val,_ := strconv.ParseFloat(fields[0],64)
	return val
}


func meminfo() {
	file, _ := os.Open("/proc/meminfo")
	data, _ := io.ReadAll(file)

	var total, used, free, buffcache, available float64
	var swaptotal, swapfree, swapused float64


	total = getvolume(data, 0)/1024
	free = getvolume(data, 1)/1024
	available = getvolume(data, 2)/1024
	used = total - available
	buffcache = getvolume(data, 3)/1024 + getvolume(data, 4)/1024

	swaptotal = getvolume(data, 14)/1024
	swapfree = getvolume(data, 15)/1024
	swapused =  swaptotal - swapfree
 
	fmt.Printf("MiB Mem : %8.0f total %8.0f free %8.0f used %8.0f buff/cache \n", total, free, used, buffcache)
	fmt.Printf("MiB Swap: %8.0f total %8.0f free %8.0f used %8.0f avail Mem \n",swaptotal, swapfree, swapused, available )
}


func main() {
	meminfo()
}
