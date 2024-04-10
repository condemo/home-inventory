package data

import (
	"context"
	"database/sql"

	"github.com/condemo/home-inventory/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

type SqliteStore struct {
	db bun.DB
}

func createSqliteStore() (*sql.DB, error) {
	sqldb, err := sql.Open(
		sqliteshim.ShimName, "file:inventory.db?cache=shared",
	)
	if err != nil {
		return nil, err
	}

	return sqldb, nil
}

func initSqliteDB() (*SqliteStore, error) {
	db, err := createSqliteStore()
	if err != nil {
		return nil, err
	}
	bunDB := bun.NewDB(db, sqlitedialect.New())

	// Tables
	_, err = bunDB.NewCreateTable().Model((*models.Cacharro)(nil)).
		IfNotExists().Exec(context.Background())

	return &SqliteStore{db: *bunDB}, err
}

func (s *SqliteStore) SaveItem(item *models.Cacharro) error {
	_, err := s.db.NewInsert().Model(item).Exec(context.Background())
	return err
}
