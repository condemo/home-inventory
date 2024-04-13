package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/condemo/home-inventory/screens"
)

type currentView int

const (
	mainView currentView = iota
	addPlaceView
)

type Model struct {
	viewList []tea.Model
	focused  currentView
	quitting bool
}

func NewModel() tea.Model {
	m := Model{}
	m.viewList = append(m.viewList, screens.NewMainModel())
	m.focused = mainView

	return m
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.viewList[m.focused].Update(msg)
	}
	return m, cmd
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}
	s := m.viewList[m.focused].View()
	return s
}

func main() {
	m := NewModel()
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
