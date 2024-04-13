package main

import "fmt"

func main() {
	program := tea.NewProgram(initialModel(), tea.WithAltScreen())

	if _, err := program.Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
}
}
