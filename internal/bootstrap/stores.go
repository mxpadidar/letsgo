package bootstrap

import (
	"github.com/jmoiron/sqlx"
	"github.com/mxpadidar/letsgo/internal/core/services"
	"github.com/mxpadidar/letsgo/internal/core/stores"
	"github.com/mxpadidar/letsgo/internal/infrastructure/dbstores"
)

type Store struct {
	Users   stores.UserStore
	Permits stores.PermitStore
}

func NewStore(db *sqlx.DB, logger services.LogService) *Store {
	return &Store{
		Users:   dbstores.NewUserDBStore(db, logger),
		Permits: dbstores.NewPermitDBStore(db, logger),
	}
}
