package system

import (
	"os"
	"strconv"
	"strings"
)

type TasksState struct {
	Total, Running, Sleeping, Stopped, Zombie int16
}

type ProcessingInitState struct {
	PID, RES, SHR          int32
	PR, NI, VIRT           int16
	USER, S, TIME, COMMAND string
	CPU, MEM               float32
}

func Processing() (TasksState, []ProcessingInitState) {

	dir, _ := os.ReadDir("/proc")

	var tasks TasksState = TasksState{0, 0, 0, 0, 0}
	var list = []ProcessingInitState{}

	for _, e := range dir {
		if !e.IsDir() {
			continue
		}

		if _, err := strconv.Atoi(e.Name()); err != nil {
			continue
		}

		pid := e.Name()

		tasks = Processingtasks(tasks, pid)
		ProcessingItem := ProcessingList(pid)
		list = append(list, ProcessingItem)
	}

	return tasks, list
}

func Processingtasks(tasks TasksState, pid string) TasksState {

	data, err := os.ReadFile("/proc/" + pid + "/stat")
	if err != nil {
		return tasks
	}

	fields := strings.Fields(string(data))
	state := fields[2]

	switch state {
	case "R":
		tasks.Running += 1
	case "S", "D":
		tasks.Sleeping += 1
	case "T":
		tasks.Stopped += 1
	case "Z":
		tasks.Zombie += 1
	}

	tasks.Total++

	return tasks
}

func ProcessingList(pid string) ProcessingInitState {
	return ProcessingInitState{}
}
