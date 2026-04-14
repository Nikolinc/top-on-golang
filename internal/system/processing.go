package system

import (
	"os"
	"strconv"
	"top-on-golang/internal/system/tasks"
)

type ProcessingInitState struct {
	PID, RES, SHR          int32
	PR, NI, VIRT           int16
	USER, S, TIME, COMMAND string
	CPU, MEM               float32
}

func Processing(totalMem float32) (tasks.TaskTotalState, []ProcessingInitState) {

	dir, _ := os.ReadDir("/proc")

	var totalstate tasks.TaskTotalState = tasks.TaskTotalState{0, 0, 0, 0, 0}
	var list = []ProcessingInitState{}

	for _, e := range dir {
		if !e.IsDir() {
			continue
		}

		if _, err := strconv.Atoi(e.Name()); err != nil {
			continue
		}

		pid := e.Name()

		totalstate = tasks.TotalState(totalstate, pid)
		ProcessingItem := ProcessingList(pid, totalMem)
		list = append(list, ProcessingItem)
	}

	return totalstate, list
}


func ProcessingList(pid string, totalMem float32) ProcessingInitState {

	unit := ProcessingInitState{}
	pidInt, _ := strconv.ParseInt(pid, 10, 32)

	unit.PID = int32(pidInt)
	unit.USER = tasks.GetUserFromPID(pid)
	unit.PR, unit.NI , unit.S = tasks.StatTask(pid)
	unit.VIRT, unit.RES, unit.SHR, unit.MEM = tasks.StatmTask(pid, totalMem)

	return unit
}
