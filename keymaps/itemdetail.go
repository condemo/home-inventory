package keymaps

import "github.com/charmbracelet/bubbles/key"

type ItemDetailKeyMap struct {
	Back key.Binding
	Del  key.Binding
	Mod  key.Binding
	Quit key.Binding
}

func (k ItemDetailKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Back, k.Mod, k.Del, k.Quit,
	}
}

func (k ItemDetailKeyMap) FullHelp() [][]key.Binding { return nil }

var ItemDetailKeys = ItemDetailKeyMap{
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "go back"),
	),
	Mod: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "modify"),
	),
	Del: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "delete"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}
