package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func showHelper() string {
	return "-------------------Usage-------------------\n\n - Press 's' to start work session.\n          s %m start work session for %m minutes\n - Press 'b' to take break.\n          b %m to take break for %m minutes\n - Press 'l' to list all completed today's sessions.\n          l YYYY-MM-DD to list completed sessions on that date.\n - Press 'q' to quit.\n"
}

func checkValidMinute(m *model, command string) (int, bool) {
	if command == "s" || command == "b" {
		return 0, true
	}

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

func loadSessions() []session {
	data, err := os.ReadFile("db.json")
	if err != nil {
		return []session{}
	}

	sessions := []session{}
	err = json.Unmarshal(data, &sessions)
	if err != nil {
		return []session{}
	}

	return sessions
}

func saveSessions(sessions []session) error {
	data, err := json.MarshalIndent(sessions, "", "  ")
	if err != nil {
		return fmt.Errorf("Error marshal data: %v", err.Error())
	}

	err = os.WriteFile("db.json", data, 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err.Error())
	}

	return nil
}

func printSessions(sessions []session, differentDate bool, date time.Time) {
	if !differentDate {
		today := time.Now()
		todaySessions := getCorrectSession(sessions, today)

		fmt.Println("Today's Completed Sessions:")
		printHelper(todaySessions)
	} else {
		differentDateSessions := getCorrectSession(sessions, date)

		fmt.Printf("Completed sessions on %v:\n", date.Format(time.DateOnly))
		printHelper(differentDateSessions)
	}
}

func getCorrectSession(sessions []session, date time.Time) []session {
	resultSessions := []session{}
	for _, s := range sessions {
		if s.StartTime.Add(s.Duration).Format(time.DateOnly) == date.Format(time.DateOnly) {
			resultSessions = append(resultSessions, s)
		}
	}

	return resultSessions
}

func printHelper(sessions []session) {
	for _, s := range sessions {
		fmt.Printf("Pomodoro session: duration of %f minutes from %v to %v\n", s.Duration.Minutes(), s.StartTime.Format("15:04"), s.EndTime.Format("15:04"))
	}
}
