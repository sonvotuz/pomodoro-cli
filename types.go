package main

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/timer"
)

type model struct {
	percent       float64
	progress      progress.Model
	timer         timer.Model
	currentTimer  *time.Timer
	timerDuration time.Duration
	remainingTime time.Duration
	startTime     time.Time
	inSession     bool
	sessionType   string // "Work" or "Break"
	keys          keyMap
	interrupting  bool
}

type keyMap struct {
	Start key.Binding
	Break key.Binding
	List  key.Binding
	Stop  key.Binding
	Quit  key.Binding
}
