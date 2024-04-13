package main

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
)

type model struct {
	percent       float64
	progress      progress.Model
	currentTimer  *time.Timer
	timerDuration time.Duration
	remainingTime time.Duration
	startTime     time.Time
	inSession     bool
	sessionType   string // "Work" or "Break"
}

func initialModel() model {
	return model{}
}
