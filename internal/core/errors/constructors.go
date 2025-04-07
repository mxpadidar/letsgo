package errors

import "fmt"

// formatErrMsg formats an error message with the given arguments.
// TODO: handle invalid arguments, or errors on formatting, gracefully.
func formatErrMsg(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

// constructor function for each error type
func NewAuthErr(format string, args ...interface{}) error {
	if format == "" {
		return NewAppError("authentication failed", ErrAuthFailed)
	}
	return NewAppError(formatErrMsg(format, args...), ErrAuthFailed)
}

func NewAccessErr(format string, args ...interface{}) error {
	if format == "" {
		return NewAppError("access denied", ErrAccessDenied)
	}
	return NewAppError(formatErrMsg(format, args...), ErrAccessDenied)
}

func NewValidationErr(format string, args ...interface{}) error {
	if format == "" {
		return NewAppError("validation failed", ErrValidation)
	}
	return NewAppError(formatErrMsg(format, args...), ErrValidation)
}

func NewInternalErr(format string, args ...interface{}) error {
	if format == "" {
		return NewAppError("internal error", ErrInternal)
	}
	return NewAppError(formatErrMsg(format, args...), ErrInternal)
}

func NewNotFoundErr(format string, args ...interface{}) error {
	if format == "" {
		return NewAppError("not found", ErrNotFound)
	}
	return NewAppError(formatErrMsg(format, args...), ErrNotFound)
}

func NewConflictErr(format string, args ...interface{}) error {
	if format == "" {
		return NewAppError("conflict", ErrConflict)
	}
	return NewAppError(formatErrMsg(format, args...), ErrConflict)
}
