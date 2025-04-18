package errors

import (
	"errors"
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

	ErrUnknown ErrorType = "unknown"
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

func GetErrorType(err error) ErrorType {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.ErrType
	}
	return ErrUnknown
}
