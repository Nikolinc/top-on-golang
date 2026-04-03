package ui

import (
	"fmt"
)

func ShowProceing() {
	fmt.Printf(" \n")
	fmt.Printf("\033[97;40mPID USER  PR  NI  VIRT  RES  SHR S  %CPU  %MEM  TIME+ COMMAND\033[0m \n")
}
