package screens

import (
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/condemo/home-inventory/models"
)

type AddPlaceView struct {
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
	return &AddPlaceView{nameEntry: input}
}

func (m AddPlaceView) Init() tea.Cmd {
	return nil
}

func (m AddPlaceView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.nameEntry.Focused() {
				m.CreatePlace()
				return ModelList[MainView].Update(nil)
			}
		case "esc":
			return ModelList[MainView].Update(nil)
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
	return m.nameEntry.View()
}

func (m AddPlaceView) CreatePlace() {
	p := new(models.Place)
	p.Name = m.nameEntry.Value()

	err := store.SavePlace(p)
	if err != nil {
		log.Panic("error:", err)
	}
}
