package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/condemo/home-inventory/screens"
)

func main() {
	screens.ModelList = []tea.Model{screens.New(), screens.NewPlaceModel()}
	m := screens.ModelList[screens.MainView]
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
