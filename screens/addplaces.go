package screens

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/condemo/home-inventory/elements"
	"github.com/condemo/home-inventory/keymaps"
	"github.com/condemo/home-inventory/models"
	"github.com/condemo/home-inventory/styles"
)

type AddPlaceView struct {
	err       error
	help      help.Model
	keys      keymaps.PlacesKeymaps
	nameEntry textinput.Model
	placeID   int64
	update    bool
	quitting  bool
}

func NewPlaceModel() *AddPlaceView {
	input := textinput.New()
	input.Prompt = " $ "
	input.Placeholder = "Insert Name..."
	input.CharLimit = 100
	input.Width = 100
	input.Validate = validateLetters
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
				if m.update {
					m.updatePlace(&models.Place{ID: m.placeID, Name: m.nameEntry.Value()})
					if m.err == nil {
						return ModelList[SelectPlace].Update(DBUpdated(true))
					}
				} else {
					p, err := m.CreatePlace()
					if err != nil {
						m.err = err
						return m, cmd
					} else {
						return ModelList[SelectPlace].Update(p)
					}
				}
			}
		case key.Matches(msg, m.keys.Back):
			return ModelList[SelectPlace].Update(nil)

		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit

		}

	case *models.Place:
		m.update = true
		m.nameEntry.SetValue(msg.Name)
		m.nameEntry.CursorEnd()
		// m.nameEntry.Focus()
		m.placeID = msg.ID
	}

	if m.nameEntry.Focused() {
		m.nameEntry, cmd = m.nameEntry.Update(msg)
		return m, cmd
	}

	return m, cmd
}

func (m AddPlaceView) View() string {
	err := ""
	if m.quitting {
		return ""
	}

	if m.err != nil {
		err = elements.NewErrorView(m.err)
	}

	entry := styles.MainInputContainer.Render(m.nameEntry.View())

	m.help.Styles.ShortKey = styles.HelpStyle
	helpView := m.help.View(m.keys)

	return lipgloss.JoinVertical(
		lipgloss.Center, entry, err, helpView)
}

func (m *AddPlaceView) CreatePlace() (*models.Place, error) {
	p := new(models.Place)

	if m.nameEntry.Value() == "" {
		return nil, fmt.Errorf("error: empty place name")
	}

	p.Name = m.nameEntry.Value()
	err := store.SavePlace(p)
	if err != nil {
		return nil, fmt.Errorf("database error: %s", err)
	}
	m.err = nil

	return p, nil
}

func (m *AddPlaceView) updatePlace(pl *models.Place) {
	if pl.Name == "" {
		m.err = fmt.Errorf("error: %s", "empty name entry")
		return
	}
	err := store.UpdatePlace(pl)
	if err != nil {
		m.err = fmt.Errorf("database error: %s", err.Error())
	}
}
