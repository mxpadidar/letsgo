package types

type Permission int64

const (
	PermUserRead Permission = iota << 1
	PermUserWrite
	PermUserDelete
)

const PermUserAll = PermUserRead | PermUserWrite | PermUserDelete
