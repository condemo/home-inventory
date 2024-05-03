package utils

import (
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/condemo/home-inventory/models"
)

func TableRowToItem(tr table.Row) *models.Cacharro {
	id, _ := strconv.ParseInt(tr[0], 10, 64)
	amount, _ := strconv.Atoi(tr[2])

	return &models.Cacharro{
		ID:     id,
		Name:   tr[1],
		Amount: uint8(amount),
		Tags:   tr[4],
	}
}
