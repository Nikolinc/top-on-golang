package tasks

import (
	"os"
	"strings"
	"top-on-golang/internal/utils"
)

func StatTask(pid string)(int16,int16, string){

		data, _ := os.ReadFile("/proc/stat")
		fields := strings.Fields(string(data))
		
		pr := utils.ParsInt16(fields[17])
		ni := utils.ParsInt16(fields[18])
		state := fields[2]

	return pr, ni, state
}

