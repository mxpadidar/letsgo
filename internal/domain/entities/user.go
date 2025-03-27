package entities

import (
	"time"

	"github.com/mxpadidar/letsgo/internal/domain/enums"
)

type User struct {
	ID             int            `json:"id" db:"id"`
	Username       string         `json:"username" db:"username"`
	HashedPassword string         `json:"-" db:"hashed_password"`
	Role           enums.UserRole `json:"role" db:"role"`
	CreatedAt      time.Time      `json:"created_at" db:"created_at"`
}

func NewUser(username string, hashedPassword string, role enums.UserRole) *User {
	return &User{
		Username:       username,
		HashedPassword: hashedPassword,
		Role:           role,
		CreatedAt:      time.Now(),
	}
}

func (u *User) AddRole(role enums.UserRole) {
	u.Role = u.Role | role
}

func (u *User) RemoveRole(role enums.UserRole) {
	u.Role = u.Role &^ role
}

func (u *User) HasRole(role enums.UserRole) bool {
	return u.Role&role == role
}
