// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package test

import (
	"database/sql"
)

type Carro struct {
	ID      int32
	Marca   string
	Modelo  string
	Preco   string
	Created sql.NullTime
}
