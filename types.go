package main

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
)

type model struct {
	percent       float64
	progress      progress.Model
	timerDuration time.Duration
	remainingTime time.Duration
	startTime     time.Time
	inSession     bool
	sessionType   string // "Work" or "Break"
	keys          keyMap
	opening       bool
	closing       bool
}

type keyMap struct {
	Start key.Binding
	Break key.Binding
	List  key.Binding
	Stop  key.Binding
	Quit  key.Binding
}
