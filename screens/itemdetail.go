package screens

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/condemo/home-inventory/keymaps"
)

type ItemDetailView struct {
	keys     keymaps.ItemDetailKeyMap
	quitting bool
}

func NewItemDetailView() *ItemDetailView {
	return &ItemDetailView{
		keys: keymaps.ItemDetailKeys,
	}
}

func (m ItemDetailView) Init() tea.Cmd { return nil }

func (m ItemDetailView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit
		}
	}
	return m, cmd
}

func (m ItemDetailView) View() string {
	if m.quitting {
		return ""
	}
	return "Working"
}
