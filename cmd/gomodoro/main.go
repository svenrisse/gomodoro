package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	application "github.com/svenrisse/gomodoro/internal"
)

func main() {
	app := tea.NewProgram(application.InitialApp(), tea.WithAltScreen())
	_, err := app.Run()
	if err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}
