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
	_, err := s.db.NewInsert().Model(item).
		Exec(context.Background())

	return err
}

func (s *SqliteStore) GetItem(id int64) (*models.Cacharro, error) {
	item := new(models.Cacharro)
	err := s.db.NewSelect().Model(item).
		Where("c.id = ?", id).
		Relation("Place").Where("place_id = place.id").
		Scan(context.TODO())

	return item, err
}

func (s *SqliteStore) GetAllItems() ([]models.Cacharro, error) {
	var il []models.Cacharro
	err := s.db.NewSelect().Model(&il).
		Relation("Place").Where("place_id = place.id").
		Scan(context.TODO())

	return il, err
}

func (s *SqliteStore) DeleteItem(id int64) error {
	it := new(models.Cacharro)
	_, err := s.db.NewDelete().Model(it).Where("id = ?", id).Exec(context.Background())

	return err
}

func (s *SqliteStore) SavePlace(item *models.Place) error {
	_, err := s.db.NewInsert().Model(item).Exec(context.TODO())
	return err
}

func (s *SqliteStore) GetAllPlaces() ([]models.Place, error) {
	var pl []models.Place
	err := s.db.NewSelect().Model(&pl).
		Order("id ASC").
		Scan(context.TODO())

	return pl, err
}

func (s *SqliteStore) GetPlace(id int64) (*models.Place, error) {
	p := new(models.Place)
	err := s.db.NewSelect().Model(p).Where("id = ?", id).Scan(context.Background())

	return p, err
}
