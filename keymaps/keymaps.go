package keymaps

import "github.com/charmbracelet/bubbles/key"

type MainKeyMap struct {
	Up      key.Binding
	Down    key.Binding
	AddItem key.Binding
	Plus    key.Binding
	Minus   key.Binding
	Back    key.Binding
	Help    key.Binding
	Search  key.Binding
	Select  key.Binding
	Quit    key.Binding
}

func (k MainKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Back, k.Help, k.Quit}
}

func (k MainKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Select},
		{k.AddItem, k.Minus, k.Plus},
		{k.Search, k.Help, k.Quit},
	}
}

var MainKeys = MainKeyMap{
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
	Plus: key.NewBinding(
		key.WithKeys("+"),
		key.WithHelp("+", "+1 can"),
	),
	Minus: key.NewBinding(
		key.WithKeys("-"),
		key.WithHelp("-", "-1 can"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
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
