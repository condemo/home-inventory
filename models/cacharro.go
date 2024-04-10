package models

import "github.com/uptrace/bun"

type Cacharro struct {
	bun.BaseModel `bun:"table:cacharro,alias:c"`

	Name    string `bun:"name,notnull,unique"`
	Place   *Place `bun:"rel:belongs-to,join:place_id=id"`
	Tags    string `bun:"tags,notnull"`
	ID      int64  `bun:",pk,autoincrement"`
	PlaceID int64  `bun:",notnull"`
	Amount  uint8  `bun:"amount,notnull"`
}

type Place struct {
	bun.BaseModel `bun:"table:place,alias:p"`

	Name string `bun:"name,unique"`
	ID   int64  `bun:",pk,autoincrement"`
}
