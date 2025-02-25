package types

import "errors"

var ErrNotFound = errors.New("not found")

var ErrInvalidCredentials = errors.New("invalid credentials")

var ErrConflict = errors.New("conflict")

var ErrValidation = errors.New("validation error")
