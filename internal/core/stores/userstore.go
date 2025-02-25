package stores

import (
	"context"

	"github.com/mxpadidar/letsgo/internal/core/entities"
)

type UserStore interface {
	FindByUsername(ctx context.Context, username string) (*entities.User, error)
	Save(ctx context.Context, user *entities.User) error
}
