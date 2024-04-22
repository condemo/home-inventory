package keymaps

import "github.com/charmbracelet/bubbles/key"

type SelectPlaceKeymap struct {
	Back   key.Binding
	Select key.Binding
	Add    key.Binding
	Quit   key.Binding
}

func (k SelectPlaceKeymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Back, k.Select, k.Add, k.Quit}
}

func (k SelectPlaceKeymap) FullHelp() [][]key.Binding {
	return nil
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
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}
