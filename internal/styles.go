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
	return lipgloss.JoinVertical(lipgloss.Left, title, fmt.Sprintf("%d", value))
}
