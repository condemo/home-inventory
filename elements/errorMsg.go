package elements

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/condemo/home-inventory/styles"
)

func NewErrorView(err error) string {
	styledErr := lipgloss.NewStyle().
		Foreground(styles.Colors.TextPrimary).Render(err.Error())
	return styles.ErrorContainer.Render(styledErr)
}
