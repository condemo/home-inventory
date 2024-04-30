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

var Colors = ColorPalette{
	SelectPrimary: lipgloss.CompleteColor{
		TrueColor: "#5f00ff",
		ANSI256:   "57",
		ANSI:      "4",
	},
	ErrPrimary: lipgloss.CompleteColor{
		TrueColor: "#af0000",
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

var (
	HelpStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(Colors.HelpPrimary)

	ContainerStyle = lipgloss.NewStyle().
			Padding(1).MarginLeft(3).BorderStyle(lipgloss.RoundedBorder())

	HelpContainer = HelpStyle.Copy().
			MarginTop(2).Align(lipgloss.Center)
	CenterContainer = lipgloss.NewStyle().Align(lipgloss.Center)

	PlacesContainer        = lipgloss.NewStyle()
	SelectedContainerStyle = ContainerStyle.Copy().
				BorderForeground(lipgloss.Color("205"))

	ErrorContainer = lipgloss.NewStyle().
			Align(lipgloss.Center).BorderForeground(Colors.ErrPrimary).
			BorderBackground(Colors.ErrPrimary).Width(40).MarginTop(1).
			BorderStyle(lipgloss.RoundedBorder()).Background(Colors.ErrPrimary)

	MainInputContainer = CenterContainer.Copy().
				Padding(1, 0).Margin(2).BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(Colors.SelectPrimary)

	InputContainer = CenterContainer.Copy().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("57"))

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
