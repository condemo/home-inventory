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

	// TABLES
	// Cacharro
	_, err = bunDB.NewCreateTable().Model((*models.Cacharro)(nil)).
		IfNotExists().Exec(context.Background())
	if err != nil {
		return nil, err
	}

	// Place
	_, err = bunDB.NewCreateTable().Model((*models.Place)(nil)).
		IfNotExists().Exec(context.Background())

	return &SqliteStore{db: *bunDB}, err
}

func (s *SqliteStore) SaveItem(item *models.Cacharro) error {
	_, err := s.db.NewInsert().Model(item).Exec(context.Background())
	return err
}

func (s *SqliteStore) SavePlace(item *models.Place) error {
	_, err := s.db.NewInsert().Model(item).Exec(context.TODO())
	return err
}

func (s *SqliteStore) GetAllPlaces() ([]models.Place, error) {
	var pl []models.Place
	err := s.db.NewSelect().Model(&pl).Scan(context.TODO())

	return pl, err
}

func (s *SqliteStore) GetPlace(id int64) (*models.Place, error) {
	p := new(models.Place)
	err := s.db.NewSelect().Model(p).Where("id = ?", id).Scan(context.Background())

	return p, err
}
