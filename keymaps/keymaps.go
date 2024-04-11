package keymaps

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up       key.Binding
	Down     key.Binding
	AddItem  key.Binding
	AddPlace key.Binding
	Help     key.Binding
	Search   key.Binding
	Select   key.Binding
	Quit     key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Select, k.Search},
		{k.AddItem, k.AddPlace},
		{k.Help, k.Quit},
	}
}

var AppKeys = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "move down"),
	),
	AddItem: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "add item"),
	),
	AddPlace: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "add place"),
	),
	Search: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "search"),
	),
	Select: key.NewBinding(
		key.WithKeys(" ", "enter"),
		key.WithHelp("space/enter", "select"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q/ctrl+c", "quit"),
	),
}
