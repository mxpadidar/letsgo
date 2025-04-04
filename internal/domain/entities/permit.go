package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/mxpadidar/letsgo/internal/domain/types"
)

type Permit struct {
	ID       uuid.UUID  `json:"id" db:"id"`
	UserID   int        `json:"user_id" db:"user_id"`
	Role     types.Role `json:"role" db:"role"`
	IssuedAt time.Time  `json:"issued_at" db:"issued_at"`
}

func NewPermit(id uuid.UUID, userID int, role types.Role, issuedAt time.Time) *Permit {
	return &Permit{ID: id, UserID: userID, Role: role, IssuedAt: issuedAt}
}
