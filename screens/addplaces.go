package screens

import (
	"log"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/condemo/home-inventory/keymaps"
	"github.com/condemo/home-inventory/models"
)

type AddPlaceView struct {
	help      help.Model
	keys      keymaps.PlacesKeymaps
	nameEntry textinput.Model
	quitting  bool
}

func NewPlaceModel() *AddPlaceView {
	input := textinput.New()
	input.Prompt = " $ "
	input.Placeholder = "Insert Name..."
	input.CharLimit = 100
	input.Width = 100
	input.Focus()
	return &AddPlaceView{
		nameEntry: input, keys: keymaps.PlacesKeys, help: help.New(),
	}
}

func (m AddPlaceView) Init() tea.Cmd {
	return nil
}

func (m AddPlaceView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Submit):
			if m.nameEntry.Focused() {
				m.CreatePlace()
				return ModelList[ItemView].Update(nil)
			}
		case key.Matches(msg, m.keys.Back):
			return ModelList[ItemView].Update(nil)
		}
	}
	if m.nameEntry.Focused() {
		m.nameEntry, cmd = m.nameEntry.Update(msg)
		return m, cmd
	}
	return m, cmd
}

func (m AddPlaceView) View() string {
	if m.quitting {
		return ""
	}

	helpView := m.help.View(m.keys)
	return lipgloss.JoinVertical(
		lipgloss.Center, m.nameEntry.View(), helpView)
}

func (m AddPlaceView) CreatePlace() {
	p := new(models.Place)
	p.Name = m.nameEntry.Value()

	err := store.SavePlace(p)
	if err != nil {
		log.Panic("error:", err)
	}
}
