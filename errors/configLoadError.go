package errors

import (
	"fmt"
)

// ConfigLoadError is for errors during loading and reading Parser`s config
type ConfigLoadError struct {
	Message    string
	InnerError error
}

func (e ConfigLoadError) Error() string {
	return fmt.Sprintf("Could not load config: %s", e.Message)
}

// NewConfigLoadError creates ConfigLoadError
func NewConfigLoadError(message string, inner error) ConfigLoadError {
	return ConfigLoadError{Message: message, InnerError: inner}
}
