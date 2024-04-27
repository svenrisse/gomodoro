package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/svenrisse/gomodoro/internal/models"
)

func main() {
	app := tea.NewProgram(models.InitialApp())
	_, err := app.Run()
	if err != nil {
		fmt.Printf("There's ben an error: %v", err)
		os.Exit(1)
	}
}
