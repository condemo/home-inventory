package styles

import "github.com/charmbracelet/lipgloss"

var (
	selectPrimary = lipgloss.CompleteColor{
		TrueColor: "#5f00ff",
		ANSI256:   "57",
		ANSI:      "4",
	}

	textPrimary = lipgloss.CompleteColor{
		TrueColor: "#ffffaf",
		ANSI256:   "229",
		ANSI:      "11",
	}

	helpPrimary = lipgloss.CompleteColor{
		TrueColor: "#afafaf",
		ANSI256:   "145",
		ANSI:      "7",
	}
)

var (
	HelpStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(helpPrimary)

	ContainerStyle = lipgloss.NewStyle().
			Padding(1)

	SelectedStyle = lipgloss.NewStyle().
			Foreground(textPrimary).
			Background(selectPrimary)
)
