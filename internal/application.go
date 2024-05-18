package application

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

type settingState uint

const (
	durationSetting settingState = iota
	pomoCountSetting
	shortBreakSetting
	longBreakSetting
)

type application struct {
	Start            time.Time
	Duration         time.Duration
	ShortBreak       time.Duration
	LongBreak        time.Duration
	Count            uint8
	PomoCountChoices uint8
	State            string
	activeModel      settingState
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
		activeModel:      durationSetting,
		Help:             help.New(),
		Keymap:           NewKeymap(),
		Progress:         progress.New(progress.WithSolidFill("#ff5f00")),
	}
}

func (app application) Init() tea.Cmd {
	return nil
}

func (app application) View() string {
	s := AppTitle("🍅 Gomodoro Timer 🍅")
	if app.State == "settings" {
		s += lipgloss.NewStyle().
			Width(72).Height(1).Align(lipgloss.Center).Bold(true).
			Render("\nSettings:\n")
		s += "\n"
		if app.activeModel == durationSetting {
			s += DurationFocus(app)
		}
		if app.activeModel == pomoCountSetting {
			s += CountFocus(app)
		}
		if app.activeModel == shortBreakSetting {
			s += ShortFocus(app)
		}
		if app.activeModel == longBreakSetting {
			s += LongFocus(app)
		}
		s += app.Keymap.helpView(app.Help)
	}

	if app.State == "focus" {
		s += fmt.Sprintf("\n%s\n", app.Timer.View())
		s += app.Progress.View()
		s += app.Keymap.focusView(app.Help)
	}

	if app.State == "shortBreak" {
		s += "Short break\n"
		s += fmt.Sprintf("%s\n", app.Timer.View())
		s += app.Progress.View()
		s += app.Keymap.focusView(app.Help)
	}

	if app.State == "longBreak" {
		s += "Long break\n"
		s += fmt.Sprintf("%s\n", app.Timer.View())
		s += app.Progress.View()
		s += app.Keymap.focusView(app.Help)
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
		if app.State == "focus" {
			app.Progress.SetPercent(0)
			if app.Count <= app.PomoCountChoices {
				app.Timer = timer.New(time.Minute * time.Duration(app.LongBreak))
				app.State = "longBreak"
			} else {
				app.Count++
				app.Timer = timer.New(time.Minute * time.Duration(app.ShortBreak))
				app.State = "shortBreak"
			}
			return app, app.Timer.Init()
		}
		if app.State == "longBreak" || app.State == "shortBreak" {
			app.Progress.SetPercent(0)
			app.Timer = timer.New(time.Minute * time.Duration(app.Duration))
			app.State = "focus"
			if app.State == "longBreak" {
				app.Count = 0
			}
			return app, app.Timer.Init()
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
				if app.activeModel == durationSetting {
					if app.Duration < 60 {
						app.Duration++
					}
				} else if app.activeModel == pomoCountSetting {
					if app.PomoCountChoices < 11 {
						app.PomoCountChoices++
					}
				} else if app.activeModel == shortBreakSetting {
					if app.ShortBreak < 60 {
						app.ShortBreak++
					}
				} else if app.activeModel == longBreakSetting {
					if app.LongBreak < 120 {
						app.LongBreak++
					}
				}
			case key.Matches(msg, app.Keymap.down):
				if app.activeModel == durationSetting {
					if app.Duration > 1 {
						app.Duration--
					}
				} else if app.activeModel == pomoCountSetting {
					if app.PomoCountChoices > 2 {
						app.PomoCountChoices--
					}
				} else if app.activeModel == shortBreakSetting {
					if app.ShortBreak > 1 {
						app.ShortBreak--
					}
				} else if app.activeModel == longBreakSetting {
					if app.LongBreak > 1 {
						app.LongBreak--
					}
				}
			case key.Matches(msg, app.Keymap.right):
				if app.activeModel == durationSetting {
					app.activeModel = pomoCountSetting
				} else if app.activeModel == pomoCountSetting {
					app.activeModel = shortBreakSetting
				} else if app.activeModel == shortBreakSetting {
					app.activeModel = longBreakSetting
				}
			case key.Matches(msg, app.Keymap.left):
				if app.activeModel == longBreakSetting {
					app.activeModel = shortBreakSetting
				} else if app.activeModel == shortBreakSetting {
					app.activeModel = pomoCountSetting
				} else if app.activeModel == pomoCountSetting {
					app.activeModel = durationSetting
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
