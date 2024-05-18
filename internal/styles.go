package application

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	UnfocusedModelStyle = lipgloss.NewStyle().
				Width(16).
				Height(4).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("188"))
	FocusedModelStyle = lipgloss.NewStyle().
				Width(16).
				Height(4).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("202"))
)

func SettingsCardContent(title string, value any) string {
	return lipgloss.JoinVertical(lipgloss.Left, title, fmt.Sprintf("%s", value))
}

func AppTitle(title string) string {
	return lipgloss.NewStyle().
		Width(72).Height(1).Align(lipgloss.Center).Foreground(lipgloss.Color("202")).
		Render(fmt.Sprintf("%v\n", title))
}
