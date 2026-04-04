package ui

import(
	"fmt"
	"top-on-golang/internal/system"
)

// tline - terminal line, строка терминала для вывода информации
func ShowCPU(cpu system.CPUState, tline int) {
	fmt.Printf("\033[%d;0H", tline)
	fmt.Printf(
		"Cpu(s): %4.1f us, %4.1f sy, %4.1f ni, %4.1f id, %4.1f wa, %4.1f hi, %4.1f si,  %4.1f st \n", 
		cpu.User,  cpu.System, cpu.Nice, cpu.Idle,
		cpu.Iowait, cpu.Irq, cpu.Softirq, cpu.Steal,
	)
}