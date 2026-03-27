package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func meminfo() {
	file, _ := os.Open("/proc/meminfo")
	data, _ := io.ReadAll(file)
	line := strings.Split(string(data), "\n")[0]
	fields := strings.Fields(line)[1:]

	fmt.Printf(fields[0])
}

func main() {
	meminfo()
}
