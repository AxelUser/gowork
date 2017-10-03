package errors

import (
	"fmt"
)

// LoadSkillError is for errors during loading skills
type LoadSkillError struct {
	Skill      string
	Message    string
	InnerError error
}

func (e LoadSkillError) Error() string {
	return fmt.Sprintf("Could not load skill '%s': %s", e.Skill, e.Message)
}

// NewLoadSkillError creates LoadSkillError
func NewLoadSkillError(skill string, message string, inner error) LoadSkillError {
	return LoadSkillError{Skill: skill, Message: message, InnerError: inner}
}
