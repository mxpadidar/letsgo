package dbstores

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/mxpadidar/letsgo/internal/core/entities"
	"github.com/mxpadidar/letsgo/internal/core/errors"
	"github.com/mxpadidar/letsgo/internal/core/services"
	"github.com/mxpadidar/letsgo/internal/core/types"
)

type UserDBStore struct {
	db     *sqlx.DB
	logger services.LogService
}

func NewUserDBStore(db *sqlx.DB, logger services.LogService) *UserDBStore {
	return &UserDBStore{db: db, logger: logger}
}

func (s *UserDBStore) Persist(ctx context.Context, user *entities.User) error {
	query := `INSERT INTO auth.users (username, hashed_password)
			VALUES ($1, $2) RETURNING id, created_at`

	row := s.db.QueryRowxContext(ctx, query, user.Username, user.HashedPassword)

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
			return nil, errors.NewNotFoundErr("user with id %d not found", id)
		}
		s.logger.Errorf("unexpected error in UserDBStore.GetByID: %v", err)
		return nil, err
	}

	return &user, nil
}

func (s *UserDBStore) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User

	query := "SELECT * FROM auth.users WHERE username = $1"

	if err := s.db.GetContext(ctx, &user, query, username); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundErr("user with username `%s` not found", username)
		}

		s.logger.Errorf("unexpected error in UserDBStore.GetByUsername: %v", err)
		return nil, err
	}

	return &user, nil
}

func (s *UserDBStore) List(ctx context.Context, paginate *types.Paginate) ([]*entities.User, error) {
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
		return nil, errors.NewValidationErr("invalid order parameter: %s", paginate.Order)
	}

	// execute query
	query := fmt.Sprintf("SELECT * FROM auth.users ORDER BY %s %s LIMIT %d OFFSET %d",
		paginate.Order,
		paginate.Direction,
		paginate.Limit,
		paginate.Offset,
	)

	if err := s.db.SelectContext(ctx, &users, query); err != nil {
		s.logger.Errorf("unexpected error in UserDBStore.List: %v", err)
		return nil, err
	}

	return users, nil
}
