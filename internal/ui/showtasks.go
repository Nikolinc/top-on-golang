package ui

import (
	"fmt"
	"top-on-golang/internal/system"
)

func ShowTasks(tasks system.TasksState) {
	fmt.Printf("Tasks: %4d total, %4d running, %4d sleeping, %4d stopped, %4d zombie \n",
		tasks.Total, tasks.Running, tasks.Sleeping, tasks.Stopped, tasks.Zombie,
	)
}
