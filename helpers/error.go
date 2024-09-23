package helpers

import "fmt"

type UniqueViolationError struct {
	Field      string
	StatusCode int
}

func (e *UniqueViolationError) Error() string {
	return fmt.Sprintf("%s must be unique", e.Field)
}

type ValidationError struct {
	Message    string
	StatusCode int
}

func (e *ValidationError) Error() string {
	return e.Message
}
