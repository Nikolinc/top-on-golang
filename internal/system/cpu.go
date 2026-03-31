package system

import (
	"io"
	"os"
	"strings"
	"strconv"
	"time"
)

type CPUState struct{
	Total, User, Nice, System, Idle, Iowait, Irq, Softirq, Steal float64
}

func ReadStat() CPUState {

	file, _ := os.Open("/proc/stat")
	defer file.Close()
	data, _ := io.ReadAll(file)
	line := strings.Split(string(data), "\n")[0]
	fields := strings.Fields(line)[1:]

	user, _ := strconv.ParseFloat(fields[0], 64)
	nice, _ := strconv.ParseFloat(fields[1], 64)
	system, _ := strconv.ParseFloat(fields[2], 64)
	idle, _ := strconv.ParseFloat(fields[3], 64)
	iowait, _ := strconv.ParseFloat(fields[4], 64)
	irq, _ := strconv.ParseFloat(fields[5], 64)
	softirq, _ := strconv.ParseFloat(fields[6], 64)
	steal, _ := strconv.ParseFloat(fields[7], 64)
	
	total:= user + nice + system + idle + iowait + 
					irq + softirq + steal

	return CPUState{
		Total:total,
		User: user, Nice: nice, System: system, Idle:idle,
		Iowait: iowait, Irq: irq, Softirq: softirq, Steal: steal,
	}
}


func CPU() CPUState{

	firshStateCPU := ReadStat()
	time.Sleep(time.Second)
	secondStateCPU := ReadStat()

	deltaUser := secondStateCPU.User - firshStateCPU.User
	deltaNice := secondStateCPU.Nice - firshStateCPU.Nice
	deltaSystem := secondStateCPU.System - firshStateCPU.System
	deltaIdle := secondStateCPU.Idle - firshStateCPU.Idle
	deltaIowait := secondStateCPU.Iowait - firshStateCPU.Iowait
	deltaIrq := secondStateCPU.Irq - firshStateCPU.Irq
	deltaSoftirq := secondStateCPU.Softirq - firshStateCPU.Softirq
	deltaSteal := secondStateCPU.Steal - firshStateCPU.Steal

	total := deltaUser + deltaNice + deltaSystem + deltaIdle +
         deltaIowait + deltaIrq + deltaSoftirq + deltaSteal

	return CPUState{
		User:    deltaUser / total * 100,
		Nice:    deltaNice / total * 100,
		System:  deltaSystem / total * 100,
		Idle:    deltaIdle / total * 100,
		Iowait:  deltaIowait / total * 100,
		Irq:     deltaIrq / total * 100,
		Softirq: deltaSoftirq / total * 100,
		Steal:   deltaSteal / total * 100,
	}
}