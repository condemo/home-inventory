package models

import (
	"fmt"

	"github.com/uptrace/bun"
)

type Cacharro struct {
	bun.BaseModel `bun:"table:cacharros,alias:c"`

	Name    string `bun:"name,notnull,unique"`
	Place   *Place `bun:"rel:belongs-to,join:place_id=id"`
	Tags    string `bun:"tags,notnull"`
	ID      int64  `bun:",pk,autoincrement"`
	PlaceID int64  `bun:",notnull"`
	Amount  uint8  `bun:"amount,notnull"`
}

type Place struct {
	bun.BaseModel `bun:"table:places,alias:p"`

	Name string `bun:"name,unique"`
	ID   int64  `bun:",pk,autoincrement"`
}

func (p Place) Title() string       { return fmt.Sprintf("%v", p.ID) }
func (p Place) Description() string { return p.Name }
func (p Place) FilterValue() string { return p.Name }
