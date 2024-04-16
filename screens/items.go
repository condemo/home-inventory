package screens

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/condemo/home-inventory/keymaps"
	"github.com/condemo/home-inventory/styles"
)

type inputsType int

const (
	nameInput inputsType = iota
	amountInput
	placeInput
	tagsInput
)

type AddItemsView struct {
	help       help.Model
	keys       keymaps.ItemsKeymaps
	inputs     []textinput.Model
	focusIndex inputsType
	quitting   bool
}

func NewItemsView() *AddItemsView {
	m := AddItemsView{
		help:       help.New(),
		keys:       keymaps.ItemsKeys,
		focusIndex: nameInput,
		inputs:     make([]textinput.Model, 4),
	}

	for i := range m.inputs {
		t := textinput.New()
		t.Cursor.Style = styles.CursorSelectStyle
		switch i {
		case int(nameInput):
			t.Placeholder = "nombre"
			t.CharLimit = 25
			t.Width = 30
			t.PromptStyle = styles.InputFocusedStyle
			t.TextStyle = styles.TextPrimaryStyle
			t.Focus()
		case int(amountInput):
			t.Placeholder = "can."
			t.CharLimit = 4
			t.Width = 4
		case int(placeInput):
			t.Placeholder = "lugar"
			t.CharLimit = 100
			t.Width = 100
		case int(tagsInput):
			t.Placeholder = "tags"
			t.CharLimit = 100
			t.Width = 100
		}
		m.inputs[i] = t
	}

	return &m
}

func (m AddItemsView) Init() tea.Cmd { return nil }

func (m AddItemsView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		// TODO: Cambiar de estilo el moverse por los elementos
		case key.Matches(msg, m.keys.Up):
			m.inputs[m.focusIndex].Blur()
			if m.focusIndex == nameInput {
				m.focusIndex = tagsInput
			} else {
				m.focusIndex--
			}
			m.inputs[m.focusIndex].Focus()

		case key.Matches(msg, m.keys.Down, m.keys.Submit):
			m.inputs[m.focusIndex].Blur()
			if m.focusIndex == tagsInput {
				m.focusIndex = nameInput
				if key.Matches(msg, m.keys.Submit) {
					// TODO: Implementar guardado de items, debería haber uno/s botones al final
					m.quitting = true
					return m, tea.Quit
				}
			} else {
				m.focusIndex++
			}
			m.inputs[m.focusIndex].Focus()

		case key.Matches(msg, m.keys.Back):
			return ModelList[MainView].Update(nil)

		case key.Matches(msg, m.keys.AddPlace):
			return ModelList[PlaceView].Update(nil)

		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit
		}
	}

	cmd = m.updateInputs(msg)
	return m, cmd
}

func (m *AddItemsView) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
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

	helpView := m.help.View(m.keys)

	mainContainer := lipgloss.JoinVertical(
		lipgloss.Left,
		b.String(),
		styles.HelpContainer.Render(helpView))

	return styles.ContainerStyle.Render(mainContainer)
}
