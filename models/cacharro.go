package models

import "github.com/uptrace/bun"

type Place string

type Cacharro struct {
	bun.BaseModel `bun:"table:cacharro,alias:c"`
	Name          string `bun:"name,notnull,unique"`
	Place         Place  `bun:"place,unique"`
	Tags          string `bun:"tags,notnull"`
	ID            int64  `bun:",pk,autoincrement"`
	Amount        uint8  `bun:"amount,notnull"`
}
