package pgstores

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/mxpadidar/letsgo/internal/core/entities"
	"github.com/mxpadidar/letsgo/internal/core/types"
)

type UserPgStore struct {
	db *sqlx.DB
}

func NewUserPgStore(db *sqlx.DB) *UserPgStore {
	return &UserPgStore{db}
}

func (store UserPgStore) FindById(ctx context.Context, id int) (*entities.User, error) {
	var user entities.User

	query := "SELECT * FROM accounts.users WHERE id = $1"

	err := store.db.GetContext(ctx, &user, query, id)
	if err == sql.ErrNoRows {
		return nil, types.ErrResourceNotFound
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (store UserPgStore) FindByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User

	query := "SELECT * FROM accounts.users WHERE username = $1"

	err := store.db.GetContext(ctx, &user, query, username)
	if err == sql.ErrNoRows {
		return nil, types.ErrResourceNotFound
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (store UserPgStore) Save(ctx context.Context, user *entities.User) error {
	query := `INSERT INTO accounts.users
		(username, hash_password, fname, lname, is_admin, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	row := store.db.QueryRowxContext(
		ctx,
		query,
		user.Username,
		user.HashPassword,
		user.FName,
		user.LName,
		user.IsAdmin,
		user.CreatedAt,
	)

	if err := row.Scan(&user.ID); err != nil {
		return types.ErrValidation
	}

	return nil
}
