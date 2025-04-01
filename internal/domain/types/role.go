package types

type Role int

const (
	RoleAdmin Role = 1 << iota
	RoleMember
	RoleGuest
)
