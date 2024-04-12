package elements

import (
	"log"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/condemo/home-inventory/data"
	"github.com/condemo/home-inventory/styles"
)

func NewTable(store data.Store) table.Model {
	colums := []table.Column{
		{Title: "Nombre", Width: 20},
		{Title: "Can.", Width: 4},
		{Title: "Lugar", Width: 30},
		{Title: "Tags", Width: 25},
	}

	// TODO: Limpiar y buscar lugar definitivo para cargar la DB
	itemsList, err := store.GetAllItems()
	if err != nil {
		log.Panic(err)
	}
	rows := []table.Row{}
	for i := range itemsList {
		current := itemsList[i]
		rows = append(rows, table.Row{
			current.Name, strconv.Itoa(int(current.Amount)), current.Place.Name, current.Tags,
		},
		)
	}

	t := table.New(
		table.WithColumns(colums),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.ThickBorder()).
		BorderForeground(styles.Colors.SelectPrimary).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(styles.Colors.TextPrimary).
		Background(styles.Colors.SelectPrimary).
		Bold(false)

	t.SetStyles(s)

	return t
}
