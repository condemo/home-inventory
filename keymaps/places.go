package keymaps

import (
	"github.com/charmbracelet/bubbles/key"
)

type PlacesKeymaps struct {
	Back   key.Binding
	Submit key.Binding
}

func (k PlacesKeymaps) ShortHelp() []key.Binding {
	return []key.Binding{k.Back, k.Submit}
}

func (k PlacesKeymaps) FullHelp() [][]key.Binding {
	return nil
}

var PlacesKeys = PlacesKeymaps{
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "go back"),
	),
	Submit: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "create"),
	),
}
