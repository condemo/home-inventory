package screens

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/condemo/home-inventory/elements"
	"github.com/condemo/home-inventory/keymaps"
	"github.com/condemo/home-inventory/models"
	"github.com/condemo/home-inventory/styles"
)

var (
	focusedButton = styles.InputFocusedStyle.Copy().
			Background(styles.Colors.SelectPrimary).
			Foreground(styles.Colors.TextPrimary)
	blurredButton = styles.BlurredStyle
)

const (
	inName int = iota
	inAmount
	inPlace
	inTags
)

type AddItemsView struct {
	err        error
	help       help.Model
	keys       keymaps.ItemsKeymaps
	inputs     []textinput.Model
	focusIndex int
	placeID    int64
	itemID     int64
	update     bool
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
			t.CharLimit = 30
			t.Width = 25
			t.PromptStyle = styles.InputFocusedStyle
			t.TextStyle = styles.TextPrimaryStyle
			t.Validate = validateLetters
			t.Focus()
		case inAmount:
			t.Placeholder = "can."
			t.CharLimit = 4
			t.Width = 4
			t.Validate = validateAmount
		case inPlace:
			t.Prompt = "Lugar: "
			t.SetValue("[ Select ]")
			t.CharLimit = 40
			t.Width = 25
			t.Validate = validateLetters
		case inTags:
			t.Placeholder = "tags"
			t.CharLimit = 80
			t.Width = 80
			t.Validate = validateLetters
		}
		m.inputs[i] = t
	}

	return &m
}

func (m AddItemsView) Init() tea.Cmd {
	return nil
}

func (m *AddItemsView) inputsToItem() *models.Cacharro {
	a, _ := strconv.ParseUint(m.inputs[inAmount].Value(), 10, 8)

	item := &models.Cacharro{
		ID:      m.itemID,
		Name:    m.inputs[inName].Value(),
		Amount:  uint8(a),
		PlaceID: m.placeID,
		Tags:    m.inputs[inTags].Value(),
	}

	return item
}

func (m *AddItemsView) Next() {
	if m.focusIndex == len(m.inputs) {
		m.focusIndex = inName
	} else {
		m.focusIndex++
	}
}

func (m *AddItemsView) Previous() {
	if m.focusIndex == inName {
		m.focusIndex = len(m.inputs)
	} else {
		m.focusIndex--
	}
}

func (m *AddItemsView) createItem() {
	item := m.inputsToItem()
	if item.Name == "" {
		m.err = fmt.Errorf("error: empty name input")
		m.focusIndex = inName - 1
		return
	}

	err := store.SaveItem(item)
	if err != nil {
		m.err = fmt.Errorf("error: database error")
		m.focusIndex = inName - 1
		return
	}
	m.err = nil
}

func (m *AddItemsView) UpdateItem() {
	item := m.inputsToItem()

	err := store.UpdateItem(item)
	if err != nil {
		m.err = fmt.Errorf("error: database error")
		m.focusIndex = inName - 1
		return
	}
	m.err = nil
}

func (m AddItemsView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Up, m.keys.Down, m.keys.Submit):
			if key.Matches(msg, m.keys.Submit) && m.focusIndex == len(m.inputs) {
				var itemStatus DBUpdated
				if m.update {
					m.UpdateItem()
					itemStatus = true
					return ModelList[MainView].Update(itemStatus)
				} else {
					m.createItem()
					if m.err == nil {
						itemStatus = true
						return ModelList[MainView].Update(itemStatus)
					}
				}
			}
			if key.Matches(msg, m.keys.Up) {
				m.Previous()
			}
			if key.Matches(msg, m.keys.Down, m.keys.Submit) {
				if key.Matches(msg, m.keys.Submit) && m.focusIndex == inPlace {
					ModelList[ItemView] = m
					return ModelList[SelectPlace].Update(WindowH)
				}
				if key.Matches(msg, m.keys.Down) {
					m.Next()
				} else {
					m.focusIndex++
				}
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
			return ModelList[MainView].Update(DBUpdated(true))

		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit
		}

	case *models.Place:
		m.placeID = msg.ID
		m.inputs[inPlace].SetValue(msg.Name)
		m.inputs[inPlace].Blur()
		m.focusIndex++
		m.inputs[inTags].Focus()

	case table.Row:
		itemID, _ := strconv.ParseInt(msg[0], 10, 64)
		m.inputs[inName].SetValue(msg[1])
		m.inputs[inAmount].SetValue(msg[2])
		m.inputs[inPlace].SetValue(msg[3])
		m.inputs[inTags].SetValue(msg[4])
		m.itemID = itemID
		m.update = true
		m.inputs[inName].Focus()
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

func (m AddItemsView) View() string {
	var b strings.Builder

	if m.quitting {
		return ""
	}

	inCont := lipgloss.JoinVertical(
		lipgloss.Left,
		styles.InputContainer.Render(m.inputs[inName].View()),
		styles.InputContainer.Render(m.inputs[inAmount].View()),
		styles.InputContainer.Render(m.inputs[inPlace].View()),
		styles.InputContainer.Render(m.inputs[inTags].View()),
	)

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}

	var btnStr string
	if m.update {
		btnStr = button.Render("[ Actualizar ]")
	} else {
		btnStr = button.Render("[ Crear ]")
	}

	fmt.Fprintf(&b, "%s", btnStr)

	if m.err != nil {
		b.WriteString("\n")
		b.WriteString(elements.NewErrorView(m.err))
	}

	m.help.Styles.ShortKey = styles.HelpStyle
	helpView := m.help.View(m.keys)

	mainContainer := lipgloss.JoinVertical(
		lipgloss.Center,
		inCont,
		b.String(),
		styles.HelpContainer.Render(helpView))

	return styles.ContainerStyle.Render(mainContainer)
}

func validateAmount(s string) error {
	_, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return err
	}
	return nil
}

func validateLetters(s string) error {
	errStr := "error invalid character"

	for _, r := range s {
		if unicode.IsDigit(r) {
			return nil
		}
		if unicode.IsSpace(r) {
			return nil
		}
		if !unicode.IsLetter(r) {
			return fmt.Errorf(errStr)
		}
	}
	return nil
}
