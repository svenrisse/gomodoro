package models

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type application struct {
	Start            time.Time
	Duration         int
	ShortBreak       time.Duration
	LongBreak        time.Duration
	PomoCountChoices []uint8
	ChoicesSet       bool
	timer            timer.Model
}

func InitialApp() application {
	return application{
		Start:      time.Now(),
		Duration:   25,
		ChoicesSet: false,
	}
}

func (app application) Init() tea.Cmd {
	return nil
}

func (app application) View() string {
	s := ""
	if !app.ChoicesSet {
		s += "How long should one Pomodoro be?\n"
		s += fmt.Sprintf("%d", app.Duration)
	}

	if app.ChoicesSet {
		s += app.timer.View()
	}

	s += "\nPress q to quit.\n"

	return s
}

func (app application) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return app, tea.Quit
		}

	case timer.TickMsg:
		var cmd tea.Cmd
		app.timer, cmd = app.timer.Update(msg)
		return app, cmd
	}

	if app.ChoicesSet == false {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "up", "k":
				if app.Duration < 60 {
					app.Duration++
				}
			case "down", "j":
				if app.Duration > 0 {
					app.Duration--
				}
			case "enter":
				app.ChoicesSet = true
				app.timer = timer.New(time.Minute * time.Duration(app.Duration))
				return app, app.timer.Init()
			}
		}
	}

	return app, nil
}
