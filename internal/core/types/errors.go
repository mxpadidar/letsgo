package types

import "errors"

type DomainError error

var (
	ErrResourceNotFound      DomainError = errors.New("resource not found")
	ErrResourceAlreadyExists DomainError = errors.New("resource already exists")
	ErrConflict              DomainError = errors.New("conflict")
	ErrValidation            DomainError = errors.New("validation error")
	ErrForbidden             DomainError = errors.New("forbidden")
	ErrUnauthorized          DomainError = errors.New("unauthorized")
	ErrInternal              DomainError = errors.New("internal error")
)
