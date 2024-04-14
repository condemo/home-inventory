package screens

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/condemo/home-inventory/keymaps"
	"github.com/condemo/home-inventory/styles"
)

type AddItemsView struct {
	help     help.Model
	keys     keymaps.ItemsKeymaps
	quitting bool
}

func NewItemsView() *AddItemsView {
	return &AddItemsView{
		help: help.New(),
		keys: keymaps.ItemsKeys,
	}
}

func (m AddItemsView) Init() tea.Cmd { return nil }

func (m AddItemsView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Back):
			return ModelList[MainView].Update(nil)
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit
		}
	}
	return m, cmd
}

func (m AddItemsView) View() string {
	if m.quitting {
		return ""
	}
	m.help.Styles.FullKey = styles.HelpStyle
	helpView := m.help.View(m.keys)

	mainContainer := lipgloss.JoinVertical(
		lipgloss.Center, "ItemView", helpView)

	return styles.ContainerStyle.Render(mainContainer)
}
