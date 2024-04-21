package screens

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	// "github.com/condemo/home-inventory/styles"
)

type place struct {
	name string
	id   int64
}

func (p place) Title() string       { return fmt.Sprintf("%v", p.id) }
func (p place) Description() string { return p.name }
func (p place) FilterValue() string { return p.name }

type SelectPlaceView struct {
	placesList list.Model
	quitting   bool
}

func NewSelectPlaceModel() *SelectPlaceView {
	var items []list.Item
	pl, err := store.GetAllPlaces()
	if err != nil {
		log.Panic(err)
	}

	for _, p := range pl {
		items = append(items, place{
			name: p.Name,
			id:   p.ID,
		})
	}
	m := &SelectPlaceView{
		placesList: list.New(items, list.NewDefaultDelegate(), 0, 0),
	}
	m.placesList.Title = "Select a Place"

	return m
}

func (m SelectPlaceView) Init() tea.Cmd { return nil }

func (m SelectPlaceView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		}
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
