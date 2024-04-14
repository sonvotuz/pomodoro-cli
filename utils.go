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
	helpText := `
-------------------Usage-------------------

 - Press 's' to start work session.
          s <minutes> to start work session for <minutes> minutes

 - Press 'b' to take a break.
          b <minutes> to take break for <minutes> minutes

 - Press 'l' to list all completed today's sessions.
          l YYYY-MM-DD to list completed sessions on that date.

 - Press 'q' to quit.
`
	return helpText
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

func printSessions(sessions []session, differentDate bool, date time.Time) string {
	printingResult := ""

	if !differentDate {
		today := time.Now()
		todaySessions := getCorrectSession(sessions, today)

		printingResult = "Today's Completed Sessions:\n"
		printingResult += printHelper(todaySessions)
	} else {
		differentDateSessions := getCorrectSession(sessions, date)

		printingResult = fmt.Sprintf("Completed sessions on %v:\n", date.Format(time.DateOnly))
		printingResult += printHelper(differentDateSessions)
	}

	return printingResult
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

func printHelper(sessions []session) string {
	resultPrinting := ""
	if len(sessions) == 0 {
		return "\nYou haven't completed any session 😕\n"
	}
	for _, s := range sessions {
		resultPrinting += fmt.Sprintf(
			"Pomodoro session: duration of %f minutes from %v to %v\n",
			s.Duration.Minutes(),
			s.StartTime.Format("15:04"),
			s.EndTime.Format("15:04"),
		)
	}
	return resultPrinting
}
