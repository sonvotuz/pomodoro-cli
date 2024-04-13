package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 80
)

type tickMsg time.Time

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
	return model{keys: keys, timerDuration: 25 * time.Minute,
		progress: progress.New(progress.WithDefaultGradient())}
}

func (m model) Init() tea.Cmd {
	return tickCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Start):
			if !m.inSession {
				// TODO: s %m command
				m.startTime = time.Now()
				m.sessionType = "Work"
				m.timerDuration = 1 * time.Minute
				m.remainingTime = m.timerDuration
				m.inSession = true
			}
			return m, tickCmd()

		case key.Matches(msg, m.keys.Break):
			if !m.inSession {
				m.startTime = time.Now()
				m.sessionType = "Break"
				m.timerDuration = 5 * time.Minute
				m.remainingTime = m.timerDuration
				m.inSession = true
			}
			return m, tickCmd()
		case key.Matches(msg, m.keys.List):
			// TODO: printSessions func
			return m, nil
		case key.Matches(msg, m.keys.Stop):
			if m.inSession {
				m.inSession = false
			}
			return m, nil
			// case key.Matches(msg, m.keymap.reset):
			// 	m.timer.Timeout = timeout
			// case key.Matches(msg, m.keymap.start, m.keymap.stop):
			// 	return m, m.timer.Toggle()
		default:
			if !m.inSession {
				showHelper()
			}
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		fmt.Println("tick tick", m.remainingTime)
		fmt.Println("duration", m.timerDuration)
		m.remainingTime -= 1 * time.Second

		m.percent = 1 - float64(m.remainingTime.Milliseconds()-m.timerDuration.Milliseconds())/float64(m.timerDuration.Milliseconds())
		fmt.Println("percent", m.percent)
		return m, tickCmd()
	case progress.FrameMsg:
		fmt.Println("update")
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

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

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}
