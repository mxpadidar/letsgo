package errors

var (
	ConflictErr   = NewAppError("resource already exists", ErrConflict)
	AuthErr       = NewAppError("authentication failed", ErrAuthFailed)
	NotFoundErr   = NewAppError("resource not found", ErrNotFound)
	InternalErr   = NewAppError("internal server error", ErrInternal)
	AccessErr     = NewAppError("access denied", ErrAccessDenied)
	ValidationErr = NewAppError("validation failed", ErrValidation)
)
