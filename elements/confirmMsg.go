package elements

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/condemo/home-inventory/styles"
)

var (
	focusedButton = styles.InputFocusedStyle.Copy().
			Background(styles.Colors.SelectPrimary).
			Foreground(styles.Colors.TextPrimary)
	blurredButton = styles.BlurredStyle
)

type ConfirmMsg struct {
	previous tea.Model
	msg      string
	sel      bool
}

func NewYesNoMsg(s string, m tea.Model) *ConfirmMsg {
	return &ConfirmMsg{
		msg:      s,
		previous: m,
	}
}

func (m ConfirmMsg) Init() tea.Cmd { return nil }

func (m ConfirmMsg) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m.previous.Update(nil)

		case "enter":
			if m.sel {
				return m.previous.Update(true)
			} else {
				return m.previous.Update(nil)
			}

		default:
			m.sel = !m.sel
		}
	}
	return m, cmd
}

func (m ConfirmMsg) View() string {
	yes := "    si    "
	no := "    no    "
	if m.sel {
		yes = focusedButton.Render(yes)
		no = blurredButton.Render(no)
	} else {
		yes = blurredButton.Render(yes)
		no = focusedButton.Render(no)
	}

	btnCont := lipgloss.NewStyle().
		MarginTop(2).Render(lipgloss.JoinHorizontal(lipgloss.Center, yes, no))

	mainCont := styles.ContainerStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.msg, btnCont,
		),
	)
	return mainCont
}
