package errors

type ErrType int

const (
	ErrInternal ErrType = iota
	ErrValidation
	ErrNotFound
	ErrConflict
	ErrAuthFailed
	ErrPermDenied
)
