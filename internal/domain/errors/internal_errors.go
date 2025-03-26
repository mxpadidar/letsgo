package errors

import "fmt"

type ServerErr struct {
	Msg string // Human-readable error message
	Op  string // Operation/function where the error occurred
	Err error  // Original/underlying error
}

// Constructor
func NewServerErr(message string, op string, err error) *ServerErr {
	return &ServerErr{Msg: message, Op: op, Err: err}
}

// Implement error interface
func (e *ServerErr) Error() string {
	return fmt.Sprintf("Something went wrong in %s: %s", e.Op, e.Msg)
}
