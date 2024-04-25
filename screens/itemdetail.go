package screens

import (
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/condemo/home-inventory/keymaps"
	"github.com/condemo/home-inventory/styles"
)

type ItemDetailView struct {
	help     help.Model
	keys     keymaps.ItemDetailKeyMap
	item     table.Row
	titles   []string
	quitting bool
}

func NewItemDetailView() *ItemDetailView {
	return &ItemDetailView{
		keys:   keymaps.ItemDetailKeys,
		titles: []string{"ID", "Nombre", "Cantidad", "Lugar", "Tags"},
		help:   help.New(),
	}
}

func (m ItemDetailView) Init() tea.Cmd { return nil }

func (m ItemDetailView) delete() {
	id, _ := strconv.ParseInt(m.item[0], 10, 64)

	err := store.DeleteItem(id)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}

func (m ItemDetailView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit

		case key.Matches(msg, m.keys.Back):
			return ModelList[MainView].Update(nil)

		case key.Matches(msg, m.keys.Mod):
			return ModelList[ItemView].Update(m.item)

		case key.Matches(msg, m.keys.Del):
			m.delete()
			var deleted DBUpdated = true
			return ModelList[MainView].Update(deleted)
		}

	case table.Row:
		m.item = msg
	}

	return m, cmd
}

func (m ItemDetailView) View() string {
	var s string

	if m.quitting {
		return ""
	}

	for i := range m.item {
		s += styles.TextPrimaryStyle.Render(m.titles[i])
		s += "\n"
		s += m.item[i]
		s += "\n\n"
	}

	container := styles.CenterContainer.Render(s)

	helpView := styles.HelpStyle.Render(m.help.View(m.keys))

	return styles.SelectedContainerStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Center, container, helpView),
	)
}
