package main

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textarea"
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

	textarea textarea.Model
	err      string

	sessions []session
}

type keyMap struct {
	Stop key.Binding
	Quit key.Binding
}

type session struct {
	StartTime time.Time     `json:"start_time"`
	EndTime   time.Time     `json:"end_time"`
	Duration  time.Duration `json:"duration"`
}
