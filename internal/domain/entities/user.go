package entities

import (
	"time"

	"github.com/mxpadidar/letsgo/internal/domain/types"
)

type User struct {
	ID             int        `json:"id" db:"id"`
	Username       string     `json:"username" db:"username"`
	HashedPassword []byte     `json:"-" db:"hashed_password"`
	Role           types.Role `json:"role" db:"role"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
}

func NewUser(username string, hashedPassword []byte, role types.Role) *User {
	return &User{
		Username:       username,
		HashedPassword: hashedPassword,
		Role:           role,
		CreatedAt:      time.Now(),
	}
}

func (u *User) AddRole(role types.Role) {
	u.Role = u.Role | role
}

func (u *User) RemoveRole(role types.Role) {
	u.Role = u.Role &^ role
}

func (u *User) HasRole(role types.Role) bool {
	return u.Role&role == role
}
