package types

type AuthUser struct {
	Username string `json:"username"`
	Role     Role   `json:"role"`
}

func NewAuthUser(username string, role Role) *AuthUser {
	return &AuthUser{Username: username, Role: role}
}
