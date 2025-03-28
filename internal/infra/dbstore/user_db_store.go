package dbstore

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/mxpadidar/letsgo/internal/domain/dtos"
	"github.com/mxpadidar/letsgo/internal/domain/entities"
	"github.com/mxpadidar/letsgo/internal/domain/errors"
)

type UserDBStore struct {
	db *sqlx.DB
}

func NewUserDBStore(db *sqlx.DB) *UserDBStore {
	return &UserDBStore{db: db}
}

func (s *UserDBStore) Persist(ctx context.Context, user *entities.User) error {
	query := `INSERT INTO auth.users (username, hashed_password)
			VALUES ($1, $2) RETURNING id, created_at`

	row := s.db.QueryRowxContext(
		ctx,
		query,
		user.Username,
		user.HashedPassword,
	)

	if err := row.Scan(&user.ID, &user.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (s *UserDBStore) GetByID(ctx context.Context, id int) (*entities.User, error) {
	var user entities.User

	query := "SELECT * FROM auth.users WHERE id = $1"

	if err := s.db.GetContext(ctx, &user, query, id); err != nil {
		if err == sql.ErrNoRows {
			msg := fmt.Sprintf("user with id %d not found", id)
			return nil, errors.NewErr(errors.ErrNotFound, msg, err)
		}
		log.Printf("failed to get user by id: %v", err)
		return nil, err
	}

	return &user, nil
}

func (s *UserDBStore) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User

	query := "SELECT * FROM auth.users WHERE username = $1"

	if err := s.db.GetContext(ctx, &user, query, username); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewErr(errors.ErrNotFound, fmt.Sprintf("user with username `%s` not found", username), nil)
		}

		log.Printf("unexpected error in UserDBStore.GetByUsername: %v", err)
		return nil, err
	}

	return &user, nil
}

func (s *UserDBStore) List(ctx context.Context, paginate *dtos.PaginateDto) ([]*entities.User, error) {
	var users []*entities.User

	// validate order parameter
	t := reflect.TypeOf(entities.User{})
	validOrder := false
	for i := 0; i < t.NumField(); i++ {
		if strings.EqualFold(t.Field(i).Name, paginate.Order) {
			validOrder = true
			break
		}
	}

	// return error if no valid field was found
	if !validOrder {
		return nil, errors.NewErr(errors.ErrValidation, fmt.Sprintf("invalid order parameter: %s", paginate.Order), nil)
	}

	// execute query
	query := fmt.Sprintf("SELECT * FROM auth.users ORDER BY %s %s LIMIT %d OFFSET %d",
		paginate.Order,
		paginate.Direction,
		paginate.Limit,
		paginate.Offset,
	)

	if err := s.db.SelectContext(ctx, &users, query); err != nil {
		log.Printf("failed to list users: %v", err)
		return nil, err
	}

	return users, nil
}
