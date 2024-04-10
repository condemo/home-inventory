package data

import (
	"log"

	"github.com/condemo/home-inventory/models"
)

type Store interface {
	SaveItem(*models.Cacharro) error
}

func InitDatabase() Store {
	store, err := initSqliteDB()
	if err != nil {
		log.Panic("error:", err)
	}

	return store
}
