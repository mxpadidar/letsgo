package stores

import (
	"context"

	"github.com/google/uuid"
	"github.com/mxpadidar/letsgo/internal/domain/entities"
	"github.com/mxpadidar/letsgo/internal/domain/types"
)

type PermitStore interface {
	Create(ctx context.Context, userID int, role types.Role) (*entities.Permit, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Permit, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Rotate(ctx context.Context, oldPermitID uuid.UUID) (*entities.Permit, error)
}
