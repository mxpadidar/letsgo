package stores

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mxpadidar/letsgo/internal/domain/entities"
	"github.com/mxpadidar/letsgo/internal/domain/errors"
	"github.com/mxpadidar/letsgo/internal/domain/services"
	"github.com/mxpadidar/letsgo/internal/domain/types"
)

type PermitDBStore struct {
	db     *sqlx.DB
	logger services.LogService
}

func NewPermitDBStore(db *sqlx.DB, logger services.LogService) *PermitDBStore {
	return &PermitDBStore{db: db, logger: logger}
}

func (s *PermitDBStore) Create(ctx context.Context, userID int, role types.Role) (*entities.Permit, error) {
	permit := entities.NewPermit(uuid.New(), userID, role, time.Now())
	query := `
	INSERT INTO auth.permits (id, user_id, role, issued_at)
	VALUES (:id, :user_id, :role, :issued_at)`

	_, err := s.db.NamedExecContext(ctx, query, permit)
	if err != nil {
		s.logger.Errorf("failed to create permit: %v", err)
		return nil, errors.InternalErr
	}
	return permit, nil
}

func (s *PermitDBStore) GetByID(ctx context.Context, id uuid.UUID) (*entities.Permit, error) {
	var permit entities.Permit
	query := "SELECT id, user_id, role, issued_at FROM auth.permits WHERE id = :id"
	err := s.db.GetContext(ctx, &permit, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundErr("permit with id %s not found", id.String())
		}
		return nil, errors.InternalErr
	}
	return &permit, nil
}

func (s *PermitDBStore) Delete(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM auth.permits WHERE id = $1"
	result, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		s.logger.Errorf("failed to delete permit: %v", err)
		return errors.InternalErr
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		s.logger.Errorf("failed to get rows affected: %v", err)
		return errors.InternalErr
	}
	if rowsAffected == 0 {
		return errors.NewNotFoundErr("permit with id %s not found", id.String())
	}
	return nil
}

func (s *PermitDBStore) Rotate(ctx context.Context, oldPermitID uuid.UUID) (*entities.Permit, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		s.logger.Errorf("failed to begin transaction: %v", err)
		return nil, errors.InternalErr
	}

	var oldPermitUserID int
	var oldPermitRole types.Role

	delQ := "DELETE FROM auth.permits WHERE id = $1 RETURNING user_id, role"
	err = tx.QueryRowContext(ctx, delQ, oldPermitID).Scan(&oldPermitUserID, &oldPermitRole)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundErr("permit with id %s not found", oldPermitID.String())
		}
		s.logger.Errorf("failed to delete permit: %v", err)
		return nil, errors.InternalErr
	}

	newPermit := entities.NewPermit(uuid.New(), oldPermitUserID, oldPermitRole, time.Now())
	InsertQ := "INSERT INTO auth.permits (id, user_id, role, issued_at) VALUES (:id, :user_id, :role, :issued_at)"
	_, err = tx.NamedExecContext(ctx, InsertQ, newPermit)
	if err != nil {
		s.logger.Errorf("failed to insert new permit: %v", err)
		return nil, errors.InternalErr
	}

	if err := tx.Commit(); err != nil {
		s.logger.Errorf("failed to commit transaction: %v", err)
		return nil, errors.InternalErr
	}

	return newPermit, nil
}
