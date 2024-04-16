package keymaps

import "github.com/charmbracelet/bubbles/key"

type ItemsKeymaps struct {
	Up       key.Binding
	Down     key.Binding
	Back     key.Binding
	AddPlace key.Binding
	Submit   key.Binding
	Quit     key.Binding
}

func (k ItemsKeymaps) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Up, k.Down, k.AddPlace, k.Back,
		k.Submit, k.Back, k.Quit,
	}
}

func (k ItemsKeymaps) FullHelp() [][]key.Binding { return nil }

var ItemsKeys = ItemsKeymaps{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "Up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "down"),
	),
	Submit: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "next/submit"),
	),
	AddPlace: key.NewBinding(
		key.WithKeys("ctrl+a"),
		key.WithHelp("ctrl+a", "add place"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "go back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}
