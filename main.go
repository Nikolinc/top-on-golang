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

	var total, used, free, shared, buffcache, available float64

	total = getvolume(data, 0)/1024
	free = getvolume(data, 1)/1024
	available = getvolume(data, 2)/1024
	used = total - available
	shared = getvolume(data, 20)/1024
	buffcache = getvolume(data, 3)/1024 + getvolume(data, 4)/1024



	fmt.Printf("total: %.0f MB\n", total)
	fmt.Printf("used: %.0f MB\n", used)
	fmt.Printf("free: %.0f MB\n", free)
	fmt.Printf("available: %.0f MB\n", available)
	fmt.Printf("shared: %.0f MB\n", shared)
	fmt.Printf("buff/cache: %.0f MB\n", buffcache)
}


func main() {
	meminfo()
}
