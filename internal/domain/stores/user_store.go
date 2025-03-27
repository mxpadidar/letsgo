package stores

import (
	"context"

	"github.com/mxpadidar/letsgo/internal/domain/entities"
	"github.com/mxpadidar/letsgo/internal/domain/errors"
)

type UserStore interface {
	Persist(ctx context.Context, user *entities.User) *errors.Err
	List(ctx context.Context, limit, offset int, direction string) ([]*entities.User, *errors.Err)
	GetByID(ctx context.Context, id int) (*entities.User, *errors.Err)
	GetByUsername(ctx context.Context, username string) (*entities.User, *errors.Err)
}
