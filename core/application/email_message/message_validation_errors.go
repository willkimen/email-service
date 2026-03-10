package emailmessage

import (
	"fmt"
	"strings"
)

// FieldValidationError marks errors related to invalid data.
// Implementations of this interface indicate that a request failed
// due to a validation rule, not an infrastructure failure.
type FieldValidationError interface {
	error
	GetField() string
}

// EmailInvalidFormatError represents a validation error indicating
// that an email address does not match the expected format.
type EmailInvalidFormatError struct {
	text string
}

func (e *EmailInvalidFormatError) Error() string {
	return e.text
}

func (e *EmailInvalidFormatError) GetField() string {
	return ""
}

func NewEmailInvalidFormatError() error {
	return &EmailInvalidFormatError{
		text: "email format is invalid",
	}
}

// EmptyFieldError represents a validation error indicating that
// a required field was not provided.
type EmptyFieldError struct {
	text  string
	Field string
}

func (e *EmptyFieldError) Error() string {
	return e.text
}

func (e *EmptyFieldError) GetField() string {
	return e.Field
}

func NewEmptyFieldError(field string) error {
	text := fmt.Sprintf("%s field is required", strings.ToLower(field))
	return &EmptyFieldError{
		text:  text,
		Field: field,
	}
}
