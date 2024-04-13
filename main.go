package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	program := tea.NewProgram(initialModel(), tea.WithAltScreen())

	if _, err := program.Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
}
