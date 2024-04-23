package screens

import (
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/condemo/home-inventory/keymaps"
	"github.com/condemo/home-inventory/models"
)

type SelectPlaceView struct {
	placesList list.Model
	keys       keymaps.SelectPlaceKeymap
	quitting   bool
}

func NewSelectPlaceModel() *SelectPlaceView {
	var items []list.Item
	pl, err := store.GetAllPlaces()
	if err != nil {
		log.Panic(err)
	}

	for _, p := range pl {
		items = append(items, &p)
	}

	m := &SelectPlaceView{
		placesList: list.New(items, list.NewDefaultDelegate(), 0, 0),
	}
	m.placesList.Title = "Select a Place"
	m.keys = keymaps.SelectPlKeymap
	m.placesList.DisableQuitKeybindings()
	m.placesList.AdditionalShortHelpKeys = keymaps.SelectPlKeymap.ShortHelp

	return m
}

func (m SelectPlaceView) Init() tea.Cmd { return nil }

func (m SelectPlaceView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit

		case key.Matches(msg, m.keys.Back):
			ModelList[SelectPlace] = m
			return ModelList[ItemView].Update(nil)

		case key.Matches(msg, m.keys.Select):
			ModelList[SelectPlace] = m
			si := m.placesList.SelectedItem().(*models.Place)
			return ModelList[ItemView].Update(si)

		case key.Matches(msg, m.keys.Add):
			ModelList[SelectPlace] = m
			return ModelList[PlaceView].Update(nil)
		}
	case *models.Place:
		m.placesList.InsertItem(len(m.placesList.Items()), msg)

	case WSize:
		m.placesList.SetHeight(int(msg) / 2)
		m.placesList.SetWidth(200)
	}

	m.placesList, cmd = m.placesList.Update(msg)
	return m, cmd
}

func (m SelectPlaceView) View() string {
	if m.quitting {
		return ""
	}
	return m.placesList.View()
}
