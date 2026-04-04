package ui

import (
	"fmt"
)

func Clear() {
	fmt.Print("\033[H\033[2J")
	// fmt.Printf("\033[H") 
}
