package application

import "github.com/charmbracelet/lipgloss"

func DurationFocus(app application) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		FocusedModelStyle.Render(
			SettingsCardContent("Pomo Duration", app.Duration),
		),
		UnfocusedModelStyle.Render(
			SettingsCardContent("Pomo Count", app.PomoCountChoices),
		),
		UnfocusedModelStyle.Render(
			SettingsCardContent("Short Break", app.ShortBreak),
		),
		UnfocusedModelStyle.Render(
			SettingsCardContent("Long Break", app.LongBreak),
		),
	)
}

func CountFocus(app application) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		UnfocusedModelStyle.Render(
			SettingsCardContent("Pomo Duration", app.Duration),
		),
		FocusedModelStyle.Render(
			SettingsCardContent("Pomo Count", app.PomoCountChoices),
		),
		UnfocusedModelStyle.Render(
			SettingsCardContent("Short Break", app.ShortBreak),
		),
		UnfocusedModelStyle.Render(
			SettingsCardContent("Long Break", app.LongBreak),
		),
	)
}

func ShortFocus(app application) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		UnfocusedModelStyle.Render(
			SettingsCardContent("Pomo Duration", app.Duration),
		),
		UnfocusedModelStyle.Render(
			SettingsCardContent("Pomo Count", app.PomoCountChoices),
		),
		FocusedModelStyle.Render(
			SettingsCardContent("Short Break", app.ShortBreak),
		),
		UnfocusedModelStyle.Render(
			SettingsCardContent("Long Break", app.LongBreak),
		),
	)
}

func LongFocus(app application) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		UnfocusedModelStyle.Render(
			SettingsCardContent("Pomo Duration", app.Duration),
		),
		UnfocusedModelStyle.Render(
			SettingsCardContent("Pomo Count", app.PomoCountChoices),
		),
		UnfocusedModelStyle.Render(
			SettingsCardContent("Short Break", app.ShortBreak),
		),
		FocusedModelStyle.Render(
			SettingsCardContent("Long Break", app.LongBreak),
		),
	)
}
