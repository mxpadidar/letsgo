package entities

import "time"

type User struct {
	ID           int       `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	HashPassword string    `json:"hash_password" db:"hash_password"`
	FName        string    `json:"fname" db:"fname"`
	LName        string    `json:"lname" db:"lname"`
	IsAdmin      bool      `json:"is_admin" db:"is_admin"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

func NewUser(username, hashPassword, fname, lname string, isAdmin bool) *User {
	return &User{
		Username:     username,
		HashPassword: hashPassword,
		FName:        fname,
		LName:        lname,
		IsAdmin:      isAdmin,
		CreatedAt:    time.Now(),
	}
}
