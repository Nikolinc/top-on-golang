package ui 

import{
	"fpm"
	"project/internal/system"
}

func ShowMem(meminfo system.MemStat){

	fmt.Printf("MiB Mem : %8.0f total %8.0f free %8.0f used %8.0f buff/cache \n", total, free, used, buffcache)
	fmt.Printf("MiB Swap: %8.0f total %8.0f free %8.0f used %8.0f avail Mem \n",swaptotal, swapfree, swapused, available )
}