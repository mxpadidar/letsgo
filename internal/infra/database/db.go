package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mxpadidar/letsgo/internal/domain/errors"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase(dataSourceName string) (*Database, error) {
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, errors.NewServerErr("failed to connect to database", "NewDatabase", err)
	}

	if err := db.Ping(); err != nil {
		return nil, errors.NewServerErr("failed to ping database", "NewDatabase", err)
	}

	return &Database{db: db}, nil
}

func (db *Database) Close() error {
	log.Println("closing database")
	return db.db.Close()
}
