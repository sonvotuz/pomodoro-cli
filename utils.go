package main

import (
	"strconv"
	"strings"
)

func showHelper() string {
	return "-------------------Usage-------------------\n\n - Press 's' to start work session.\n          s %m start work session for %m minutes\n - Press 'b' to take break.\n          b %m to take break for %m minutes\n - Press 'l' to list all completed today's sessions.\n - Press 'q' to quit.\n"
}

func checkValidMinute(m *model, command string) (int, bool) {
	spacing := command[1:]
	if !strings.HasPrefix(spacing, " ") {
		m.err = "Invalid command"
		return 0, false
	}

	numOfMinutesStr := strings.TrimSpace(command[2:])

	numOfMinutes, err := strconv.Atoi(numOfMinutesStr)
	if err != nil {
		m.err = "Invalid number of minutes"
		return 0, false
	}
	return numOfMinutes, true
}
