package tasks

import (
	"os"
	"os/user"
	"strings"
)

func GetUserFromPID(pid string) string {
	data, _ := os.ReadFile("/proc/" + pid + "/status")

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Uid:") {
			fields := strings.Fields(line)
			uid := fields[1]

			u, err := user.LookupId(uid)
			if err == nil {
				return u.Username
			}
		}
	}

	return "unknown"
}