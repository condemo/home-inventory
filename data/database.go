package data

import (
	"context"
	"database/sql"

	"github.com/condemo/home-inventory/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

var bunDB *bun.DB

func createDatabase() (*sql.DB, error) {
	sqldb, err := sql.Open(
		sqliteshim.ShimName, "file:inventory.db?cache=shared",
	)
	if err != nil {
		return nil, err
	}

	return sqldb, nil
}

func InitDB() error {
	db, err := createDatabase()
	if err != nil {
		return err
	}
	bunDB = bun.NewDB(db, sqlitedialect.New())

	// Tables
	_, err = bunDB.NewCreateTable().Model((*models.Cacharro)(nil)).
		IfNotExists().Exec(context.Background())

	return err
}
