package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/condemo/home-inventory/data"
	"github.com/condemo/home-inventory/keymaps"
	"github.com/condemo/home-inventory/styles"
)

type Model struct {
	help      help.Model
	keys      keymaps.KeyMap
	itemTable table.Model
	loaded    bool
	quitting  bool
}

var store = data.InitDatabase()

func NewModel() *Model {
	return &Model{
		itemTable: initTable(),
		keys:      keymaps.AppKeys,
		help:      help.New(),
	}
}

func initTable() table.Model {
	colums := []table.Column{
		{Title: "Nombre", Width: 20},
		{Title: "Can.", Width: 4},
		{Title: "Lugar", Width: 30},
		{Title: "Tags", Width: 25},
	}

	// TODO: Limpiar y buscar lugar definitivo para cargar la DB
	itemsList, err := store.GetAllItems()
	if err != nil {
		log.Panic(err)
	}
	rows := []table.Row{}
	for i := range itemsList {
		current := itemsList[i]
		rows = append(rows, table.Row{
			current.Name, strconv.Itoa(int(current.Amount)), current.Place.Name, current.Tags,
		},
		)
	}

	t := table.New(
		table.WithColumns(colums),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.ThickBorder()).
		BorderForeground(styles.Colors.SelectPrimary).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(styles.Colors.TextPrimary).
		Background(styles.Colors.SelectPrimary).
		Bold(false)

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
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
