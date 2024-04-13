package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/timer"
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

func initialModel() model {
	return model{keys: keys, timerDuration: 25 * time.Minute}
}

func (m model) Init() tea.Cmd {
	return m.timer.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
			// case key.Matches(msg, m.keymap.reset):
			// 	m.timer.Timeout = timeout
			// case key.Matches(msg, m.keymap.start, m.keymap.stop):
			// 	return m, m.timer.Toggle()
		}
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case timer.TickMsg:
		var cmds []tea.Cmd
		var cmd tea.Cmd

		m.remainingTime += m.timer.Interval
		pct := m.remainingTime.Milliseconds() * 100 / m.timerDuration.Milliseconds()
		cmds = append(cmds, m.progress.SetPercent(float64(pct)/100))

		m.timer, cmd = m.timer.Update(msg)
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)

	default:
		return m, nil
	}
}

func (m model) View() string {
	if !m.inSession {
		return showHelper()
	}
	pad := strings.Repeat(" ", padding)
	return fmt.Sprintf("\n%s Timer: %s left\n", m.sessionType, m.remainingTime) + pad + m.progress.ViewAs(m.percent) + "\n\n" + pad + helpStyle(" - Press 'x' to stop\n - Press 'q' to quit")
}
