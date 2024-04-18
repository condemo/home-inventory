package styles

import (
	"github.com/charmbracelet/lipgloss"
)

type ColorPalette struct {
	SelectPrimary  lipgloss.CompleteColor
	ErrPrimary     lipgloss.CompleteColor
	WarningPrimary lipgloss.CompleteColor
	TextPrimary    lipgloss.CompleteColor
	HelpPrimary    lipgloss.CompleteColor
}

var Colors ColorPalette

var (
	HelpStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(Colors.HelpPrimary)

	ContainerStyle = lipgloss.NewStyle().
			Padding(1).MarginLeft(3).BorderStyle(lipgloss.RoundedBorder())

	HelpContainer = lipgloss.NewStyle().
			MarginTop(3).Align(lipgloss.Center)

	PlacesContainer = lipgloss.NewStyle()

	SelectedStyle = lipgloss.NewStyle().
			Foreground(Colors.TextPrimary).
			Background(Colors.SelectPrimary)
	TextPrimaryStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("205"))

	BlurredStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	NoStyle           = lipgloss.NewStyle()
	InputFocusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("57"))
	CursorSelectStyle = InputFocusedStyle.Copy()
)

func init() {
	Colors = ColorPalette{
		SelectPrimary: lipgloss.CompleteColor{
			TrueColor: "#5f00ff",
			ANSI256:   "57",
			ANSI:      "4",
		},
		ErrPrimary: lipgloss.CompleteColor{
			TrueColor: "af0000",
			ANSI256:   "124",
			ANSI:      "1",
		},
		WarningPrimary: lipgloss.CompleteColor{
			TrueColor: "#d75f00",
			ANSI256:   "166",
			ANSI:      "11",
		},
		TextPrimary: lipgloss.CompleteColor{
			TrueColor: "#ffffaf",
			ANSI256:   "229",
			ANSI:      "11",
		},
		HelpPrimary: lipgloss.CompleteColor{
			TrueColor: "#ffffaf",
			ANSI256:   "229",
			ANSI:      "11",
		},
	}
}
