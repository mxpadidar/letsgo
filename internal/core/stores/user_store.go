package stores

import (
	"context"

	"github.com/mxpadidar/letsgo/internal/core/entities"
	"github.com/mxpadidar/letsgo/internal/core/types"
)

type UserStore interface {
	Persist(ctx context.Context, user *entities.User) error
	List(ctx context.Context, paginate *types.Paginate) ([]*entities.User, error)
	GetByID(ctx context.Context, id int) (*entities.User, error)
	GetByUsername(ctx context.Context, username string) (*entities.User, error)
}
