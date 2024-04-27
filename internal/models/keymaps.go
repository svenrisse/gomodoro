package models

import (
	"github.com/charmbracelet/bubbles/key"
)

type keymap struct {
	up      key.Binding
	down    key.Binding
	left    key.Binding
	right   key.Binding
	confirm key.Binding
	quit    key.Binding
}

func NewKeymap() keymap {
	return keymap{
		up:      key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "increase")),
		down:    key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "decrease")),
		left:    key.NewBinding(key.WithKeys("left"), key.WithHelp("←", "prev")),
		right:   key.NewBinding(key.WithKeys("right"), key.WithHelp("→", "next")),
		confirm: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "confirm")),
		quit:    key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "exit")),
	}
}
