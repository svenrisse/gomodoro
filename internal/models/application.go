package models

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 80
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

type application struct {
	Start            time.Time
	Duration         time.Duration
	ShortBreak       time.Duration
	LongBreak        time.Duration
	PomoCountChoices uint8
	ChoicesSet       bool
	timer            timer.Model
	help             help.Model
	keymap           keymap
	progress         progress.Model
}

func InitialApp() application {
	return application{
		Start:            time.Now(),
		Duration:         25,
		PomoCountChoices: 3,
		ChoicesSet:       false,
		help:             help.New(),
		keymap:           NewKeymap(),
		progress:         progress.New(progress.WithScaledGradient("#93c5fd", "#1d4ed8")),
	}
}

func (app application) Init() tea.Cmd {
	return nil
}

func (app application) View() string {
	s := "🍅 Gomodoro Timer\n\n"
	if !app.ChoicesSet {
		s += "How long should one Gomodoro be?\n"
		s += fmt.Sprintf("%d min", app.Duration)
		s += app.keymap.helpView(app.help)
	}

	if app.ChoicesSet {
		s += fmt.Sprintf("%s\n", app.timer.View())
		s += app.progress.View()
		s += app.keymap.helpView(app.help)
	}

	return s
}

func (app application) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, app.keymap.quit):
			return app, tea.Quit
		}

	case timer.TickMsg:
		if app.progress.Percent() == 1.0 {
			// pause...
		}
		var cmd tea.Cmd
		app.timer, cmd = app.timer.Update(msg)

		// TODO: incr depending on time
		progressCmd := app.progress.IncrPercent(0.1)
		return app, tea.Batch(cmd, progressCmd)

	case tea.WindowSizeMsg:
		app.progress.Width = msg.Width - padding*2 - 4
		if app.progress.Width > maxWidth {
			app.progress.Width = maxWidth
		}
		return app, nil

	case progress.FrameMsg:
		progressModel, cmd := app.progress.Update(msg)
		app.progress = progressModel.(progress.Model)
		return app, cmd

	}
	if app.ChoicesSet == false {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, app.keymap.up):
				if app.Duration < 60 {
					app.Duration++
				}
			case key.Matches(msg, app.keymap.down):
				if app.Duration > 0 {
					app.Duration--
				}
			case key.Matches(msg, app.keymap.confirm):
				app.ChoicesSet = true
				app.timer = timer.New(time.Minute * time.Duration(app.Duration))
				return app, app.timer.Init()
			}
		}
	}

	return app, nil
}
