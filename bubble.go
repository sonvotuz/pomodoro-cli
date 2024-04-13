func (m model) View() string {
	if !m.inSession {
		return showHelper()
	}
	pad := strings.Repeat(" ", padding)
	return fmt.Sprintf("\n%s Timer: %s left\n", m.sessionType, m.remainingTime) + pad + m.progress.ViewAs(m.percent) + "\n\n" + pad + helpStyle(" - Press 'x' to stop\n - Press 'q' to quit")
}
