package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/condemo/home-inventory/data"
	"github.com/condemo/home-inventory/keymaps"
	"github.com/condemo/home-inventory/models"
	"github.com/condemo/home-inventory/styles"
)

type Model struct {
	store     data.Store
	help      help.Model
	keys      keymaps.KeyMap
	itemTable table.Model
	loaded    bool
	quitting  bool
}

func NewModel() *Model {
	return &Model{
		itemTable: initTable(),
		store:     data.InitDatabase(),
		keys:      keymaps.AppKeys,
		help:      help.New(),
	}
}

func initTable() table.Model {
	colums := []table.Column{
		{Title: "Nombre", Width: 20},
		{Title: "Can.", Width: 4},
		{Title: "Lugar", Width: 15},
		{Title: "Tags", Width: 30},
	}

	// TODO: Cargar los datos dinamicamente
	rows := []table.Row{
		{"Prueba", "1", "Salón cajón 2", "video, cables"},
		{"Prueba2", "2", "Salón cajón 2", "video, cables"},
		{"Prueba3", "3", "Cama Gus", "cacharro, electrónica"},
		{"Prueba4", "53", "Salón cajón 1", "video, cables"},
		{"Prueba5", "6", "Salón cajón 2", "video, cables"},
	}

	t := table.New(
		table.WithColumns(colums),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(15),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		Align(lipgloss.Center)
	s.Selected = styles.SelectedStyle
	s.Cell = s.Cell.Align(lipgloss.Center)

	t.SetStyles(s)

	return t
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit

		case key.Matches(msg, m.keys.Down):
			if m.itemTable.Cursor() == len(m.itemTable.Rows())-1 {
				m.itemTable.GotoTop()
			} else {
				m.itemTable.SetCursor(m.itemTable.Cursor() + 1)
			}

		case key.Matches(msg, m.keys.Up):
			if m.itemTable.Cursor() == 0 {
				m.itemTable.SetCursor(len(m.itemTable.Rows()))
			} else {
				m.itemTable.SetCursor(m.itemTable.Cursor() - 1)
			}

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}
	}
	return m, cmd
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}

	m.help.Styles.FullKey = styles.HelpStyle

	helpView := m.help.View(m.keys)

	mainContainer := lipgloss.JoinVertical(
		lipgloss.Center, m.itemTable.View(), helpView)

	return styles.ContainerStyle.Render(mainContainer)
}

func main() {
	m := NewModel()
	m.store.SaveItem(&models.Cacharro{
		Name:   "Test",
		Place:  "Cajón 2",
		Tags:   "dsada dasdasd dasdasd asdaskjk",
		Amount: 2,
	})

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
