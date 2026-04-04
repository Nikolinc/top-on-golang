package ui

import (
	"fmt"
	"top-on-golang/internal/system/tasks"
)

// tline - terminal line, строка терминала для вывода информации 
func ShowTasks(tasks tasks.TaskTotalState, tline int) {
	fmt.Printf("\033[%d;0H", tline)
	fmt.Printf("Tasks: %4d total, %4d running, %4d sleeping, %4d stopped, %4d zombie \n",
		tasks.Total, tasks.Running, tasks.Sleeping, tasks.Stopped, tasks.Zombie,
	)
}