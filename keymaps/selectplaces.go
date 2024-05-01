package keymaps

import "github.com/charmbracelet/bubbles/key"

type SelectPlaceKeymap struct {
	Back   key.Binding
	Select key.Binding
	Add    key.Binding
	Delete key.Binding
	Modify key.Binding
	Quit   key.Binding
}

func (k SelectPlaceKeymap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Add, k.Select, k.Back,
	}
}

func (k SelectPlaceKeymap) FullHelp() []key.Binding {
	return []key.Binding{k.Modify, k.Delete, k.Quit}
}

var SelectPlKeymap = SelectPlaceKeymap{
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "go back"),
	),
	Select: key.NewBinding(
		key.WithKeys("enter", " "),
		key.WithHelp("enter/space", "select"),
	),
	Add: key.NewBinding(
		key.WithKeys("ctrl+a"),
		key.WithHelp("ctrl+a", "add place"),
	),
	Delete: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "delete"),
	),
	Modify: key.NewBinding(
		key.WithKeys("m"),
		key.WithHelp("m", "modify"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}
