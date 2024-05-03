package screens

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/condemo/home-inventory/data"
	"github.com/condemo/home-inventory/elements"
	"github.com/condemo/home-inventory/keymaps"
	"github.com/condemo/home-inventory/styles"
	"github.com/condemo/home-inventory/utils"
)

var ModelList []tea.Model

type (
	WSize     int
	DBUpdated bool
)

var (
	WindowW WSize
	WindowH WSize
)

type currentView int

const (
	MainView currentView = iota
	PlaceView
	ItemView
	SelectPlace
	ItemDetail
	ConfirmPopUp
)

type MainModel struct {
	help      help.Model
	keys      keymaps.MainKeyMap
	itemTable table.Model
	loaded    bool
	quitting  bool
}

var store = data.InitDatabase()

func New() *MainModel {
	return &MainModel{
		itemTable: elements.NewTable(store),
		keys:      keymaps.MainKeys,
		help:      help.New(),
	}
}

func (m MainModel) Init() tea.Cmd { return nil }

func (m *MainModel) reloadTable() {
	m.itemTable = elements.NewTable(store)
}

func (m *MainModel) changeAmount(b bool) {
	it := utils.TableRowToItem(m.itemTable.SelectedRow())
	if b {
		it.Amount += 1
	} else {
		it.Amount -= 1
	}

	store.UpdateItem(it)
	m.reloadTable()
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		WindowW = WSize(msg.Width)
		WindowH = WSize(msg.Height)
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

		case key.Matches(msg, m.keys.AddItem):
			ModelList[MainView] = m
			ModelList[ItemView] = NewItemsView()
			return ModelList[ItemView].Update(nil)

		case key.Matches(msg, m.keys.Select):
			ModelList[MainView] = m
			r := m.itemTable.SelectedRow()
			return ModelList[ItemDetail].Update(r)

		case key.Matches(msg, m.keys.Minus):
			m.changeAmount(false)

		case key.Matches(msg, m.keys.Plus):
			m.changeAmount(true)

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}

	default:
		m.reloadTable()
	}
	return m, cmd
}

func (m MainModel) View() string {
	if m.quitting {
		return ""
	}

	m.help.Styles.FullKey = styles.HelpStyle
	m.help.Styles.ShortKey = styles.HelpStyle

	helpView := m.help.View(m.keys)

	mainContainer := lipgloss.JoinVertical(
		lipgloss.Center, m.itemTable.View(), helpView)

	return styles.ContainerStyle.Render(mainContainer)
}
