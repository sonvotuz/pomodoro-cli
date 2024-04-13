package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 80
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

var keys = keyMap{
	Start: key.NewBinding(
		key.WithKeys("s"),
	),
	Break: key.NewBinding(
		key.WithKeys("b"),
	),
	List: key.NewBinding(
		key.WithKeys("l"),
	),
	Stop: key.NewBinding(
		key.WithKeys("x"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
	),
}

func (m model) Init() tea.Cmd {
	return m.timer.Init()
}
func (m model) View() string {
	if !m.inSession {
		return showHelper()
	}
	pad := strings.Repeat(" ", padding)
	return fmt.Sprintf("\n%s Timer: %s left\n", m.sessionType, m.remainingTime) + pad + m.progress.ViewAs(m.percent) + "\n\n" + pad + helpStyle(" - Press 'x' to stop\n - Press 'q' to quit")
}
