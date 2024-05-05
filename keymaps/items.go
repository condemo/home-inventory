package keymaps

import "github.com/charmbracelet/bubbles/key"

type ItemsKeymaps struct {
	Up     key.Binding
	Down   key.Binding
	Back   key.Binding
	Submit key.Binding
	Quit   key.Binding
}

func (k ItemsKeymaps) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Back, k.Up, k.Down, k.Submit, k.Quit,
	}
}

func (k ItemsKeymaps) FullHelp() [][]key.Binding { return nil }

var ItemsKeys = ItemsKeymaps{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "tab"),
		key.WithHelp("↓/tab", "down"),
	),
	Submit: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "next/submit"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}
