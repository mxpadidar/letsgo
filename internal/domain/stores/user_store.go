package stores

import (
	"context"

	"github.com/mxpadidar/letsgo/internal/domain/dtos"
	"github.com/mxpadidar/letsgo/internal/domain/entities"
)

type UserStore interface {
	Persist(ctx context.Context, user *entities.User)error
	List(ctx context.Context, paginate *dtos.PaginateDto) ([]*entities.User, error)
	GetByID(ctx context.Context, id int) (*entities.User, error)
	GetByUsername(ctx context.Context, username string) (*entities.User, error)
}
