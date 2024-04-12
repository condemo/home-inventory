package screens

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type AddPlaceView struct {
	nameEntry textinput.Model
}

func NewPlaceView() tea.Model {
	input := textinput.New()
	input.Prompt = "$"
	input.Placeholder = "Insert Name..."
	input.CharLimit = 100
	input.Width = 100
	return AddPlaceView{nameEntry: input}
}

func (m AddPlaceView) Init() tea.Cmd {
	return nil
}

func (m AddPlaceView) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m AddPlaceView) View() string {
	return "AddPlaceView"
}
