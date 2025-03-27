package enums

type UserRole int

const (
	RoleAdmin UserRole = 1 << iota
	RoleUser
)
