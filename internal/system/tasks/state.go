package tasks

import (
	"os"
	"strings"
	"strconv"
)

func ParsInt16(str string) int16 {
	v, _ := strconv.ParseInt(str, 10, 16)
	return int16(v)
}


func StatTask(pid string)(int16,int16, string){

		data, _ := os.ReadFile("/proc/stat")
		fields := strings.Fields(string(data))
		
		pr := ParsInt16(fields[17])
		ni := ParsInt16(fields[18])
		state := fields[2]

	return pr, ni, state
}

