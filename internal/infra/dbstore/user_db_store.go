package dbstore

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/mxpadidar/letsgo/internal/domain/entities"
	"github.com/mxpadidar/letsgo/internal/domain/errors"
)

type UserDBStore struct {
	db *sqlx.DB
}

func NewUserDBStore(db *sqlx.DB) *UserDBStore {
	return &UserDBStore{db: db}
}

func (s *UserDBStore) Persist(ctx context.Context, user *entities.User) *errors.Err {
	query := `INSERT INTO auth.users (username, hashed_password)
			VALUES ($1, $2) RETURNING id, created_at`

	row := s.db.QueryRowxContext(
		ctx,
		query,
		user.Username,
		user.HashedPassword,
	)

	if err := row.Scan(&user.ID, &user.CreatedAt); err != nil {
		return errors.NewErr(errors.ErrValidation, "cannot persist user", err)
	}

	return nil
}

func (s *UserDBStore) GetByID(ctx context.Context, id int) (*entities.User, *errors.Err) {
	var user entities.User

	query := "SELECT * FROM auth.users WHERE id = $1"

	if err := s.db.GetContext(ctx, &user, query, id); err != nil {
		if err == sql.ErrNoRows {
			msg := fmt.Sprintf("user with id %d not found", id)
			return nil, errors.NewErr(errors.ErrNotFound, msg, err)
		}

		return nil, errors.NewErr(errors.ErrInternal, "failed to get user by id", err)
	}

	return &user, nil
}

func (s *UserDBStore) GetByUsername(ctx context.Context, username string) (*entities.User, *errors.Err) {
	var user entities.User

	query := "SELECT * FROM auth.users WHERE username = $1"

	if err := s.db.GetContext(ctx, &user, query, username); err != nil {
		if err == sql.ErrNoRows {
			msg := fmt.Sprintf("user with username %s not found", username)
			return nil, errors.NewErr(errors.ErrNotFound, msg, err)
		}

		log.Printf("unexpected error in UserDBStore.GetByUsername: %v", err)
		return nil, errors.NewErr(errors.ErrInternal, "failed to get user by username", err)
	}

	return &user, nil
}

func (s *UserDBStore) List(ctx context.Context, limit, offset int, direction string) ([]*entities.User, *errors.Err) {
	// Implementation goes here
	return nil, nil
}
