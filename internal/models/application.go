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

type sessionState uint

const (
	durationView sessionState = iota
	pomoCountView
	shortBreakView
	longBreakView
)

var (
	unfocusedModelStyle = lipgloss.NewStyle().
				Width(15).
				Height(3).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("188"))

	focusedModelStyle = lipgloss.NewStyle().
				Width(15).
				Height(3).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("202"))
)

type application struct {
	Start            time.Time
	Duration         time.Duration
	ShortBreak       time.Duration
	LongBreak        time.Duration
	Count            uint8
	PomoCountChoices uint8
	State            string
	activeModel      sessionState
	Timer            timer.Model
	Help             help.Model
	Keymap           keymap
	Progress         progress.Model
}

func InitialApp() application {
	return application{
		Start:            time.Now(),
		Duration:         25,
		ShortBreak:       10,
		LongBreak:        15,
		Count:            0,
		PomoCountChoices: 3,
		State:            "settings",
		activeModel:      durationView,
		Help:             help.New(),
		Keymap:           NewKeymap(),
		Progress:         progress.New(progress.WithSolidFill("#0ea5e9")),
	}
}

func (app application) Init() tea.Cmd {
	return nil
}

func (app application) View() string {
	s := "🍅 Gomodoro Timer\n\n"
	if app.State == "settings" {
		s += "How long should one Gomodoro be?\n"
		if app.activeModel == durationView {
			s += lipgloss.JoinHorizontal(
				lipgloss.Top,
				focusedModelStyle.Render(fmt.Sprintf("%d min", app.Duration)),
				unfocusedModelStyle.Render(fmt.Sprintf("%d", app.PomoCountChoices)),
				unfocusedModelStyle.Render(fmt.Sprintf("%d min", app.ShortBreak)),
				unfocusedModelStyle.Render(fmt.Sprintf("%d min", app.LongBreak)),
			)
		}
		if app.activeModel == pomoCountView {
			s += lipgloss.JoinHorizontal(
				lipgloss.Top,
				unfocusedModelStyle.Render(fmt.Sprintf("%d min", app.Duration)),
				focusedModelStyle.Render(fmt.Sprintf("%d", app.PomoCountChoices)),
				unfocusedModelStyle.Render(fmt.Sprintf("%d min", app.ShortBreak)),
				unfocusedModelStyle.Render(fmt.Sprintf("%d min", app.LongBreak)),
			)
		}
		if app.activeModel == shortBreakView {
			s += lipgloss.JoinHorizontal(
				lipgloss.Top,
				unfocusedModelStyle.Render(fmt.Sprintf("%d min", app.Duration)),
				unfocusedModelStyle.Render(fmt.Sprintf("%d", app.PomoCountChoices)),
				focusedModelStyle.Render(fmt.Sprintf("%d min", app.ShortBreak)),
				unfocusedModelStyle.Render(fmt.Sprintf("%d min", app.LongBreak)),
			)
		}
		if app.activeModel == longBreakView {
			s += lipgloss.JoinHorizontal(
				lipgloss.Top,
				unfocusedModelStyle.Render(fmt.Sprintf("%d min", app.Duration)),
				unfocusedModelStyle.Render(fmt.Sprintf("%d", app.PomoCountChoices)),
				unfocusedModelStyle.Render(fmt.Sprintf("%d min", app.ShortBreak)),
				focusedModelStyle.Render(fmt.Sprintf("%d min", app.LongBreak)),
			)
		}

		s += app.Keymap.helpView(app.Help)
	}

	if app.State == "focus" {
		s += fmt.Sprintf("%s\n", app.Timer.View())
		s += app.Progress.View()
		s += app.Keymap.helpView(app.Help)
	}

	if app.State == "shortBreak" {
		s += "Short break\n"
		s += fmt.Sprintf("%s\n", app.Timer.View())
		s += app.Progress.View()
		s += app.Keymap.helpView(app.Help)
	}

	return s
}

func (app application) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, app.Keymap.quit):
			return app, tea.Quit
		}

	case timer.TickMsg:
		var cmd tea.Cmd
		app.Timer, cmd = app.Timer.Update(msg)

		var progressCmd tea.Cmd
		if app.State == "focus" {
			progressCmd = app.Progress.IncrPercent(1 / (float64(app.Duration.Abs() * 60)))
		}
		if app.State == "shortBreak" {
			progressCmd = app.Progress.IncrPercent(1 / (float64(app.ShortBreak.Abs() * 60)))
		}
		if app.State == "longBreak" {
			progressCmd = app.Progress.IncrPercent(1 / (float64(app.LongBreak.Abs() * 60)))
		}
		return app, tea.Batch(cmd, progressCmd)

	case timer.TimeoutMsg:
		// pause...
		if app.State == "focus" {
			if app.Count == app.PomoCountChoices {
				//	app.Timer = timer.New(time.Second * time.Duration(app.LongBreak))
				//	app.Count = 0
				app.State = "longPause"
			}
			if app.Count < app.PomoCountChoices {
				app.Progress.SetPercent(0)
				app.Count++
				app.Timer = timer.New(time.Minute * time.Duration(app.ShortBreak))
				app.State = "shortBreak"
			}
			return app, app.Timer.Init()
		}

		if app.State == "longPause" || app.State == "shortBreak" {
			app.State = "focus"
		}

	case tea.WindowSizeMsg:
		app.Progress.Width = msg.Width - padding*2 - 4
		if app.Progress.Width > maxWidth {
			app.Progress.Width = maxWidth
		}
		return app, nil

	case progress.FrameMsg:
		progressModel, cmd := app.Progress.Update(msg)
		app.Progress = progressModel.(progress.Model)
		return app, cmd

	}

	if app.State == "settings" {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, app.Keymap.up):
				if app.Duration < 60 {
					app.Duration++
				}
			case key.Matches(msg, app.Keymap.down):
				if app.Duration > 0 {
					app.Duration--
				}
			case key.Matches(msg, app.Keymap.right):
				if app.activeModel == durationView {
					app.activeModel = pomoCountView
				} else if app.activeModel == pomoCountView {
					app.activeModel = shortBreakView
				} else if app.activeModel == shortBreakView {
					app.activeModel = longBreakView
				}
			case key.Matches(msg, app.Keymap.left):
				if app.activeModel == longBreakView {
					app.activeModel = shortBreakView
				} else if app.activeModel == shortBreakView {
					app.activeModel = pomoCountView
				} else if app.activeModel == pomoCountView {
					app.activeModel = durationView
				}

			case key.Matches(msg, app.Keymap.confirm):
				app.State = "focus"
				app.Timer = timer.New(time.Minute * time.Duration(app.Duration))
				return app, app.Timer.Init()
			}
		}
	}

	return app, nil
}
