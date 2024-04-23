package screens

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/condemo/home-inventory/keymaps"
	"github.com/condemo/home-inventory/models"
	"github.com/condemo/home-inventory/styles"
)

var (
	focusedButton = styles.InputFocusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", styles.BlurredStyle.Render("Submit"))
)

const (
	inName int = iota
	inAmount
	inPlace
	inTags
)

type AddItemsView struct {
	help       help.Model
	keys       keymaps.ItemsKeymaps
	inputs     []textinput.Model
	focusIndex int
	placeID    int64
	quitting   bool
}

func NewItemsView() *AddItemsView {
	m := AddItemsView{
		help:       help.New(),
		keys:       keymaps.ItemsKeys,
		focusIndex: 0,
		inputs:     make([]textinput.Model, 4),
	}

	for i := range m.inputs {
		t := textinput.New()
		t.Cursor.Style = styles.CursorSelectStyle
		switch i {
		case inName:
			t.Placeholder = "nombre"
			t.CharLimit = 25
			t.Width = 30
			t.PromptStyle = styles.InputFocusedStyle
			t.TextStyle = styles.TextPrimaryStyle
			t.Focus()
		case inAmount:
			t.Placeholder = "can."
			t.CharLimit = 4
			t.Width = 4
		case inPlace:
			// t.Placeholder = "lugar"
			t.Prompt = "Place: "
			t.SetValue("[ Select ]")
			t.CharLimit = 100
			t.Width = 100
		case inTags:
			t.Placeholder = "tags"
			t.CharLimit = 100
			t.Width = 100
		}
		m.inputs[i] = t
	}

	return &m
}

func (m AddItemsView) Init() tea.Cmd {
	return nil
}

func (m AddItemsView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		// TODO: Cambiar de estilo el moverse por los elementos
		case key.Matches(msg, m.keys.Up, m.keys.Down, m.keys.Submit):
			if key.Matches(msg, m.keys.Submit) && m.focusIndex == len(m.inputs) {
				msg := m.createItem()
				ModelList[ItemView] = m
				return ModelList[MainView].Update(msg)
			}
			if key.Matches(msg, m.keys.Up) {
				if m.focusIndex == inName {
					m.focusIndex = len(m.inputs) - 1
				} else {
					m.focusIndex--
				}
			}
			if key.Matches(msg, m.keys.Down, m.keys.Submit) {
				if key.Matches(msg, m.keys.Submit) && m.focusIndex == inPlace {
					ModelList[ItemView] = m
					return ModelList[SelectPlace].Update(WindowH)
				}
				if key.Matches(msg, m.keys.Down) {
					if m.focusIndex == len(m.inputs)-1 {
						m.focusIndex = inName
					} else {
						m.focusIndex++
					}
				} else {
					m.focusIndex++
				}
			}
			if key.Matches(msg, m.keys.Down) && m.focusIndex >= len(m.inputs) {
				m.focusIndex = inName
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = styles.TextPrimaryStyle
					m.inputs[i].TextStyle = styles.TextPrimaryStyle
					continue
				}

				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = styles.NoStyle
				m.inputs[i].TextStyle = styles.NoStyle
			}

			return m, tea.Batch(cmds...)

		case key.Matches(msg, m.keys.Back):
			return ModelList[MainView].Update(nil)

		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit
		}
	case *models.Place:
		m.placeID = msg.ID
		m.inputs[inPlace].SetValue(msg.Name)
	}
	cmd = m.updateInputs(msg)

	return m, cmd
}

func (m *AddItemsView) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		if i == inPlace {
			continue
		}
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m AddItemsView) createItem() tea.Msg {
	a, err := strconv.ParseUint(m.inputs[inAmount].Value(), 10, 8)
	if err != nil {
		// TODO: Implementar notificaciones de errores y no salir de el programa
		log.Panic(err)
	}

	item := &models.Cacharro{
		Name:    m.inputs[inName].Value(),
		Amount:  uint8(a),
		PlaceID: m.placeID,
		Tags:    m.inputs[inTags].Value(),
	}

	// TODO: Manejar mejor los errores
	err = store.SaveItem(item)
	if err != nil {
		log.Panic(err)
	}
	item, err = store.GetItem(item.ID)
	if err != nil {
		log.Panic(err)
	}

	return *item
}

func (m AddItemsView) View() string {
	var b strings.Builder

	if m.quitting {
		return ""
	}

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}

	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	helpView := m.help.View(m.keys)

	mainContainer := lipgloss.JoinVertical(
		lipgloss.Left,
		b.String(),
		styles.HelpContainer.Render(helpView))

	return styles.ContainerStyle.Render(mainContainer)
}
