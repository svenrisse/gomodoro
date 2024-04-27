package models

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type application struct {
	Start            time.Time
	Duration         int8
	ShortBreak       time.Duration
	LongBreak        time.Duration
	PomoCountChoices []uint8
}

func InitialApp() application {
	return application{
		Start:            time.Now(),
		Duration:         25,
		PomoCountChoices: []uint8{1, 2, 3, 4, 5},
	}
}

func (app application) Init() tea.Cmd {
	return nil
}

func (app application) View() string {
	s := "How long should one Pomodoro be?"

	cursor := " "

	s += fmt.Sprintf("%s %d", cursor, app.Duration)

	s += "\nPress q to quit.\n"

	return s
}

func (app application) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return app, tea.Quit
		case "up", "k":
			if app.Duration < 60 {
				app.Duration++
			}
		case "down", "j":
			if app.Duration > 0 {
				app.Duration--
			}
		case "enter", " ":
			fmt.Print("enter pressed")
		}
	}

	return app, nil
}
