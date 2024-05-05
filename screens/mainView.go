package screens

import (
	"log"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
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
	searchQuery string
	help        help.Model
	searchInput textinput.Model
	keys        keymaps.MainKeyMap
	itemTable   table.Model
	loaded      bool
	quitting    bool
}

var store = data.InitDatabase()

func New() *MainModel {
	si := textinput.New()
	si.Prompt = " > "
	si.Placeholder = "Buscar"
	si.Width = 50
	si.CharLimit = 50

	k := keymaps.MainKeys
	k.Back.SetEnabled(false)

	h := help.New()

	return &MainModel{
		itemTable:   elements.NewTable(store),
		searchInput: si,
		keys:        k,
		help:        h,
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

func (m *MainModel) searchTable() {
	il, err := store.SearchItems(m.searchQuery)
	if err != nil {
		// TODO: Implementar error en esta vista
		log.Panic(err)
	}

	tr := utils.ItemsToTableRow(il)
	m.itemTable.SetRows(tr)
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
			if !m.searchInput.Focused() {
				if m.itemTable.Cursor() == len(m.itemTable.Rows())-1 {
					m.itemTable.GotoTop()
				} else {
					m.itemTable.SetCursor(m.itemTable.Cursor() + 1)
				}
			}

		case key.Matches(msg, m.keys.Up):
			if !m.searchInput.Focused() {
				if m.itemTable.Cursor() == 0 {
					m.itemTable.SetCursor(len(m.itemTable.Rows()))
				} else {
					m.itemTable.SetCursor(m.itemTable.Cursor() - 1)
				}
			}

		case key.Matches(msg, m.keys.AddItem):
			if !m.searchInput.Focused() {
				ModelList[MainView] = m
				ModelList[ItemView] = NewItemsView()
				return ModelList[ItemView].Update(nil)
			}

		case key.Matches(msg, m.keys.Select):
			if m.searchInput.Focused() {
				if msg.String() == "enter" {
					m.searchInput.Blur()
					m.searchInput.Reset()
					m.itemTable.Focus()
					if m.searchQuery == "" {
						m.keys.Back.SetEnabled(false)
						m.reloadTable()
					} else {
						m.keys.Back.SetHelp("esc", "reset")
					}
					return m, cmd
				}
			} else {
				if len(m.itemTable.Rows()) > 0 {
					ModelList[MainView] = m
					r := m.itemTable.SelectedRow()
					return ModelList[ItemDetail].Update(r)
				}
			}

		case key.Matches(msg, m.keys.Minus):
			if !m.searchInput.Focused() {
				m.changeAmount(false)
			}

		case key.Matches(msg, m.keys.Plus):
			if !m.searchInput.Focused() {
				m.changeAmount(true)
			}

		case key.Matches(msg, m.keys.Search):
			if !m.searchInput.Focused() {
				m.itemTable.SetRows([]table.Row{})
				m.itemTable.Blur()
				m.searchInput.Focus()
				m.keys.Back.SetHelp("esc", "back")
				m.keys.Back.SetEnabled(true)
				return m, cmd
			}

		case key.Matches(msg, m.keys.Back):
			m.searchQuery = ""
			m.reloadTable()
			if m.keys.Back.Enabled() {
				m.keys.Back.SetEnabled(false)
			}
			if m.searchInput.Focused() {
				m.searchInput.Blur()
				m.searchInput.Reset()
				m.itemTable.Focus()
				m.keys.Back.SetEnabled(false)

				return m, cmd
			}

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}

	case DBUpdated:
		if !m.searchInput.Focused() {
			m.reloadTable()
		}
	}

	if m.searchInput.Focused() {
		m.searchInput, cmd = m.searchInput.Update(msg)
		m.searchQuery = m.searchInput.Value()
		m.searchTable()
		return m, cmd
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

	inputView := ""
	if m.searchInput.Focused() {
		inputView = styles.InputContainer.Render(m.searchInput.View())
	}

	mainContainer := lipgloss.JoinVertical(
		lipgloss.Center, inputView, m.itemTable.View(), helpView)

	return styles.ContainerStyle.Render(mainContainer)
}
