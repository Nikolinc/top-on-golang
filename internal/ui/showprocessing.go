package ui

import (
	"fmt"
)

func ShowProceing() {
	fmt.Println(" ")
	fmt.Println("\033[97;40mPID USER  PR  NI  VIRT  RES  SHR S  %CPU  %MEM  TIME+ COMMAND\033[0m")
}
