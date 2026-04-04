package ui

import (
	"fmt"
	"top-on-golang/internal/system"
)
// tline - terminal line, строка терминала для вывода информации
func ShowMem(meminfo system.MemStat,tline int) {
	fmt.Printf("\033[%d;0H", tline)
	fmt.Printf("MiB Mem : %8.0f total %8.0f free %8.0f used %8.0f buff/cache \n",
		meminfo.Total, meminfo.Free, meminfo.Used, meminfo.BuffCache)
	fmt.Printf("\033[%d;0H", tline+1)
	fmt.Printf("MiB Swap: %8.0f total %8.0f free %8.0f used %8.0f avail Mem \n",
		meminfo.SwapTotal, meminfo.SwapFree, meminfo.SwapUsed, meminfo.Available)
}
