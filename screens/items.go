package screens

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/condemo/home-inventory/keymaps"
	"github.com/condemo/home-inventory/styles"
)

type AddItemsView struct {
	help        help.Model
	keys        keymaps.ItemsKeymaps
	nameInput   textinput.Model
	amountInput textinput.Model
	placeInput  textinput.Model
	tagsInput   textinput.Model
	quitting    bool
}

func NewItemsView() *AddItemsView {
	ni := textinput.New()
	ni.Placeholder = "nombre"
	ni.CharLimit = 25
	ni.Width = 30

	ai := textinput.New()
	ai.Placeholder = "can."
	ai.CharLimit = 4
	ai.Width = 4

	pi := textinput.New()
	pi.Placeholder = "lugar"
	pi.CharLimit = 100
	pi.Width = 100

	ti := textinput.New()
	ti.Placeholder = "tags"
	ti.CharLimit = 100
	ti.Width = 100

	ni.Focus()

	return &AddItemsView{
		help:        help.New(),
		keys:        keymaps.ItemsKeys,
		nameInput:   ni,
		amountInput: ai,
		placeInput:  pi,
		tagsInput:   ti,
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
		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit
		}
	}
	if m.nameInput.Focused() {
		m.nameInput, cmd = m.nameInput.Update(msg)
	}
	return m, cmd
}

func (m AddItemsView) View() string {
	if m.quitting {
		return ""
	}

	helpView := m.help.View(m.keys)

	mainContainer := lipgloss.JoinVertical(
		lipgloss.Left,
		m.nameInput.View(),
		m.placeInput.View(),
		m.amountInput.View(),
		m.tagsInput.View(),
		styles.HelpContainer.Render(helpView))

	return styles.ContainerStyle.Render(mainContainer)
}
