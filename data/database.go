package data

import (
	"log"

	"github.com/condemo/home-inventory/models"
)

type Store interface {
	SaveItem(*models.Cacharro) error
	SavePlace(*models.Place) error
	GetPlace(int64) (*models.Place, error)
}

func InitDatabase() Store {
	store, err := initSqliteDB()
	if err != nil {
		log.Panic("error:", err)
	}

	return store
}
