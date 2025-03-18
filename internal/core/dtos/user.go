package dtos

import (
	"time"

	"github.com/mxpadidar/letsgo/internal/core/entities"
)

type UserDto struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	FName     string    `json:"firstName"`
	LName     string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewUserDtoFromUser(user *entities.User) *UserDto {
	return &UserDto{
		ID:        user.ID,
		Username:  user.Username,
		FName:     user.FName,
		LName:     user.LName,
		CreatedAt: user.CreatedAt,
	}
}
