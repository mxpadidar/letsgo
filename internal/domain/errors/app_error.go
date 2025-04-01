package errors

import (
	"fmt"
)

type ErrorType string

func (e ErrorType) String() string {
	return string(e)
}

const (
	ErrInternal     ErrorType = "internal"
	ErrValidation   ErrorType = "validation"
	ErrNotFound     ErrorType = "not_found"
	ErrConflict     ErrorType = "conflict"
	ErrAuthFailed   ErrorType = "auth_failed"
	ErrAccessDenied ErrorType = "access_denied"
)

type AppError struct {
	Message string
	ErrType ErrorType
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%v] %s", e.ErrType, e.Message)
}

// Generic error creator
func NewAppError(msg string, errType ErrorType) *AppError {
	return &AppError{ErrType: errType, Message: msg}
}

// Convenience functions for each error type
func NewInternalError(msg string) *AppError {
	return NewAppError(msg, ErrInternal)
}

func NewValidationError(msg string) *AppError {
	return NewAppError(msg, ErrValidation)
}

func NewNotFoundError(msg string) *AppError {
	return NewAppError(msg, ErrNotFound)
}

func NewConflictError(msg string) *AppError {
	return NewAppError(msg, ErrConflict)
}

func NewAuthFailedError(msg string) *AppError {
	return NewAppError(msg, ErrAuthFailed)
}

func NewAccessDeniedError(msg string) *AppError {
	return NewAppError(msg, ErrAccessDenied)
}
