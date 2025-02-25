package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresDb struct {
	Db *sqlx.DB
}

func NewPgDb(connStr string) (*PostgresDb, error) {
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &PostgresDb{Db: db}, nil
}
