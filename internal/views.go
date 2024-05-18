package application

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func DurationFocus(app application) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		FocusedModelStyle.Render(
			SettingsCardContent("Pomo Duration", fmt.Sprintf("%d min", app.Duration)),
		),
		UnfocusedModelStyle.Render(
			SettingsCardContent("Pomo Count", fmt.Sprintf("%d", app.PomoCountChoices)),
		),
		UnfocusedModelStyle.Render(
			SettingsCardContent("Short Break", fmt.Sprintf("%d min", app.ShortBreak)),
		),
		UnfocusedModelStyle.Render(
			SettingsCardContent("Long Break", fmt.Sprintf("%d min", app.LongBreak)),
		),
	)
}

func CountFocus(app application) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		UnfocusedModelStyle.Render(
			SettingsCardContent("Pomo Duration", fmt.Sprintf("%d min", app.Duration)),
		),
		FocusedModelStyle.Render(
			SettingsCardContent("Pomo Count", fmt.Sprintf("%d", app.PomoCountChoices)),
		),
		UnfocusedModelStyle.Render(
			SettingsCardContent("Short Break", fmt.Sprintf("%d min", app.ShortBreak)),
		),
		UnfocusedModelStyle.Render(
			SettingsCardContent("Long Break", fmt.Sprintf("%d min", app.LongBreak)),
		),
	)
}

func ShortFocus(app application) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		UnfocusedModelStyle.Render(
			SettingsCardContent("Pomo Duration", fmt.Sprintf("%d min", app.Duration)),
		),
		UnfocusedModelStyle.Render(
			SettingsCardContent("Pomo Count", fmt.Sprintf("%d", app.PomoCountChoices)),
		),
		FocusedModelStyle.Render(
			SettingsCardContent("Short Break", fmt.Sprintf("%d min", app.ShortBreak)),
		),
		UnfocusedModelStyle.Render(
			SettingsCardContent("Long Break", fmt.Sprintf("%d min", app.LongBreak)),
		),
	)
}

func LongFocus(app application) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		UnfocusedModelStyle.Render(
			SettingsCardContent("Pomo Duration", fmt.Sprintf("%d min", app.Duration)),
		),
		UnfocusedModelStyle.Render(
			SettingsCardContent("Pomo Count", fmt.Sprintf("%d", app.PomoCountChoices)),
		),
		UnfocusedModelStyle.Render(
			SettingsCardContent("Short Break", fmt.Sprintf("%d min", app.ShortBreak)),
		),
		FocusedModelStyle.Render(
			SettingsCardContent("Long Break", fmt.Sprintf("%d min", app.LongBreak)),
		),
	)
}
