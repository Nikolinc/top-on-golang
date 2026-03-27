package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
	"time"
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

func loadavg() (float64, float64, float64){
	file, _ := os.Open("/proc/loadavg")
	data, _ := io.ReadAll(file)
	fields := strings.Fields(string(data))

	loadavgOne, _ := strconv.ParseFloat(fields[0], 64)
	loadavgFive, _ := strconv.ParseFloat(fields[1], 64)
	loadavgFifteen, _ := strconv.ParseFloat(fields[2], 64)


	return loadavgOne, loadavgFive, loadavgFifteen
}

func uptime(){

	file, _ := os.Open("/proc/uptime")
	data, _ := io.ReadAll(file)
	fields := strings.Fields(string(data))
	uptimeSec, _ := strconv.ParseFloat(fields[0], 64)

	var days, hours, minutes int

	days = int(uptimeSec) / 86400
	hours = (int(uptimeSec) % 86400) / 3600
	minutes = (int(uptimeSec) % 3600) / 60

	loadavgOne, loadavgFive, loadavgFifteen := loadavg()

	fmt.Printf("top - %8s up %d days, %2d:%2d 2 users, load average: %.2f %.2f %.2f\n",time.Now().Format("15:04:05"),  days, hours, minutes, loadavgOne, loadavgFive, loadavgFifteen)
	
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
	uptime()
  meminfo()
}
