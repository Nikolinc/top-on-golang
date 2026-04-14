package tasks

import (
	"os"
	"strings"
	"top-on-golang/internal/utils"
)

func StatmTask(pid string, totalMem float32) (int16, int32, int32, float32) {

	data, _ := os.ReadFile("/proc/stat")
	fields := strings.Fields(string(data))

	virt := utils.ParsInt16(fields[1])
	res := utils.ParsInt32(fields[2])
	shr := utils.ParsInt32(fields[3])
	mem := float32(shr) / totalMem * 100

	return virt, res, shr, mem
}
