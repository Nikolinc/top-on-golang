package ui

import (
	"fmt"
	"top-on-golang/internal/system"
)

// sline - start line, первая строка с которой пойдет изменения дальеше по списку
func ShowProceing(processList []system.ProcessingInitState, sline int) {

	// h - тип заголовок, да костыль но это было
	// сделано чтоб все было в одном столбце
	// и для денамического изменения заголовков в будущем

	hpid := "PID"
	huser := "USER"
	hpr := "PR"
	hni := "NI"
	hvert := "VIRT"
	hres := "RES"
	hshr := "SHR"
	hs := "S"
	hcpu := "%CPU"
	hmem := "MEM"
	htime := "TIME+"
	hcomand := "COMMAND"

	fmt.Printf(" \n")
	fmt.Printf("\033[97;40m %5s %-8s %4s %4s %-6s %-6s %-6s %-6s %-6s %-6s %s %s \033[0m\n",
		hpid, huser, hpr, hni, hvert, hres, hshr, hs, hcpu, hmem, htime, hcomand,
	)

	for tline, task := range processList {

		if tline > 30 {
			return
		}
		fmt.Printf("\033[%d;0H", tline+sline)
		fmt.Printf(" %5d %-8s %4d %4d %-6d %6d %-6d %6s %-6.0f %-6.0f \n",
			task.PID, task.USER, task.PR,
			task.NI, task.VIRT, task.RES,
			task.SHR, task.S, task.CPU,
			task.MEM)

	}

}

