package dbstore

import (
	"github.com/jmoiron/sqlx"
	"github.com/mxpadidar/letsgo/internal/domain/entities"
)

type UserDBStore struct {
	db *sqlx.DB
}

func NewUserDBStore(db *sqlx.DB) *UserDBStore {
	return &UserDBStore{db: db}
}

func (s *UserDBStore) Persist(user *entities.User) error {
	// Implementation goes here
	return nil
}

func (s *UserDBStore) GetByID(id int) (*entities.User, error) {
	// Implementation goes here
	return nil, nil
}
