package models

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type application struct {
	Start            time.Time
	Duration         time.Duration
	ShortBreak       time.Duration
	LongBreak        time.Duration
	PomoCountChoices []uint8
	ChoicesSet       bool
	timer            timer.Model
	help             help.Model
	keymap           keymap
}

func InitialApp() application {
	return application{
		Start:      time.Now(),
		Duration:   25,
		ChoicesSet: false,
		help:       help.New(),
		keymap:     NewKeymap(),
	}
}

func (app application) Init() tea.Cmd {
	return nil
}

func (app application) helpView() string {
	return "\n" + app.help.ShortHelpView([]key.Binding{
		app.keymap.up,
		app.keymap.down,
		app.keymap.left,
		app.keymap.right,
		app.keymap.confirm,
		app.keymap.quit,
	})
}

func (app application) View() string {
	s := ""
	if !app.ChoicesSet {
		s += "How long should one Pomodoro be?\n"
		s += fmt.Sprintf("%d min", app.Duration)
	}

	if app.ChoicesSet {
		s += fmt.Sprintf("🍅 %s", app.timer.View())
	}

	s += app.helpView()
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
		var cmd tea.Cmd
		app.timer, cmd = app.timer.Update(msg)
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
