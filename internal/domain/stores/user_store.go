package stores

import "github.com/mxpadidar/letsgo/internal/domain/entities"

type UserStore interface {
	Persist(user *entities.User) error
	GetByID(id int) (*entities.User, error)
}
