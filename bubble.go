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
	padding      = 2
	maxWidth     = 80
	workSession  = "Work"
	breakSession = "Break"
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
				m.sessionType = workSession
				// m.timerDuration = 25 * time.Minute
				m.timerDuration = 10 * time.Second
				m.remainingTime = m.timerDuration + 3*time.Second
				m.percent = 0
				m.inSession = true
				m.opening = true
				m.closing = false
			}
			return m, tickCmd()

		case key.Matches(msg, m.keys.Break):
			if !m.inSession {
				m.startTime = time.Now()
				m.sessionType = breakSession
				// m.timerDuration = 5 * time.Minute
				m.timerDuration = 10 * time.Second
				m.remainingTime = m.timerDuration + 3*time.Second
				m.percent = 0
				m.inSession = true
				m.opening = true
				m.closing = false
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
		if !m.inSession {
			return m, nil
		}

		m.remainingTime -= 1 * time.Second

		if m.opening {
			if m.remainingTime.Milliseconds() <= m.timerDuration.Milliseconds() {
				m.opening = false
			}

			return m, tickCmd()
		}

		if m.remainingTime.Seconds() <= -4 {
			m.closing = false
			m.inSession = false
			return m, nil
		}

		if m.remainingTime.Seconds() <= 0 {
			m.closing = true
			return m, tickCmd()
		}

		m.percent = 1 - float64(m.remainingTime.Milliseconds())/float64(m.timerDuration.Milliseconds())

		return m, tickCmd()

	default:
		return m, nil
	}
}

func (m model) View() string {
	if !m.inSession {
		return fmt.Sprintf(
			"\n%s\n%s\n\n",
			showHelper(),
			m.textarea.View(),
		)
	}
	if m.opening {
		return fmt.Sprintf("Ready to start new %s session in %d seconds...", m.sessionType, int(m.remainingTime.Seconds()-m.timerDuration.Seconds()))
	}

	if m.closing {
		if m.sessionType == workSession {
			return fmt.Sprintf("You have completed one %s session. Keep it up 💪", m.sessionType)
		}
		return fmt.Sprintf("Regained your energy with short %s. Let's start %s session.", breakSession, workSession)
	}
	pad := strings.Repeat(" ", padding)

	return fmt.Sprintf("\n%s Timer: %s left\n\n", m.sessionType, m.remainingTime) + pad + m.progress.ViewAs(m.percent) + "\n\n\n" + helpStyle(" - Press 'x' to stop\n - Press 'q' to quit")
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}
