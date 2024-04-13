package main

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 80
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

type tickMsg time.Time

func (m model) Init() tea.Cmd {
	return tickCmd()
}
func (m model) View() string {
	if !m.inSession {
		return showHelper()
	}
	pad := strings.Repeat(" ", padding)
	return fmt.Sprintf("\n%s Timer: %s left\n", m.sessionType, m.remainingTime) + pad + m.progress.ViewAs(m.percent) + "\n\n" + pad + helpStyle(" - Press 'x' to stop\n - Press 'q' to quit")
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
