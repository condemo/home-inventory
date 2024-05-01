package screens

import (
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/condemo/home-inventory/keymaps"
	"github.com/condemo/home-inventory/models"
	"github.com/condemo/home-inventory/styles"
)

type SelectPlaceView struct {
	placesList   list.Model
	keys         keymaps.SelectPlaceKeymap
	filterActive bool
	quitting     bool
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
	m.placesList.AdditionalFullHelpKeys = keymaps.SelectPlKeymap.FullHelp

	return m
}

func (m *SelectPlaceView) Reload() {
	var items []list.Item
	pl, err := store.GetAllPlaces()
	if err != nil {
		log.Panic(err)
	}
	for i := range m.placesList.Items() {
		m.placesList.RemoveItem(i)
	}

	for _, p := range pl {
		items = append(items, &p)
	}

	m.placesList.SetItems(items)
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
			if !m.filterActive {
				return ModelList[ItemView].Update(nil)
			} else {
				m.filterActive = false
			}

		case key.Matches(msg, m.keys.Select):
			ModelList[SelectPlace] = m
			if m.placesList.SelectedItem() == nil {
				return ModelList[ItemView].Update(nil)
			} else {
				si := m.placesList.SelectedItem().(*models.Place)
				return ModelList[ItemView].Update(si)
			}

		case key.Matches(msg, m.keys.Add):
			ModelList[SelectPlace] = m
			return ModelList[PlaceView].Update(nil)

		case key.Matches(msg, m.keys.Delete):
			// TODO: Impementar, si !filterActive

		case key.Matches(msg, m.keys.Modify):
			if !m.filterActive {
				ModelList[SelectPlace] = m
				si := m.placesList.SelectedItem().(*models.Place)
				return ModelList[PlaceView].Update(si)
			}

		case key.Matches(msg, m.placesList.KeyMap.Filter):
			m.filterActive = true
		}
	case *models.Place:
		m.placesList.InsertItem(len(m.placesList.Items()), msg)

	case DBUpdated:
		m.Reload()

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

	m.placesList.Help.Styles.ShortKey = styles.HelpStyle
	m.placesList.Help.Styles.FullKey = styles.HelpStyle

	return m.placesList.View()
}
