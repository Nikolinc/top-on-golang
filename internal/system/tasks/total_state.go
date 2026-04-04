package tasks

import (
	"os"
	"strings"
)

type TaskTotalState struct {
	Total, Running, Sleeping, Stopped, Zombie int16
}

func TotalState(totaltasks TaskTotalState, pid string) TaskTotalState {

	data, err := os.ReadFile("/proc/" + pid + "/stat")
	if err != nil {
		return totaltasks
	}

	fields := strings.Fields(string(data))
	state := fields[2]

	switch state {
	case "R":
		totaltasks.Running += 1
	case "S", "D":
		totaltasks.Sleeping += 1
	case "T":
		totaltasks.Stopped += 1
	case "Z":
		totaltasks.Zombie += 1
	}

	totaltasks.Total++

	return totaltasks
}

