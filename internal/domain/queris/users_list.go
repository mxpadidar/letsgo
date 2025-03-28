package queris

import (
	"context"

	"github.com/mxpadidar/letsgo/internal/domain/dtos"
	"github.com/mxpadidar/letsgo/internal/domain/entities"
	"github.com/mxpadidar/letsgo/internal/domain/stores"
)

type UsersListQuery struct {
	paginate *dtos.PaginateDto
}

func NewUsersListQuery(paginate *dtos.PaginateDto) *UsersListQuery {
	return &UsersListQuery{paginate: paginate}
}

func (q *UsersListQuery) Fetch(ctx context.Context, store stores.UserStore) ([]*entities.User, error) {
	return store.List(ctx, q.paginate)
}
