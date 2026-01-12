package emailrequest

import (
	"fmt"
	"strings"
)

type EmailInvalidFormatError struct {
	text string
}

func (e *EmailInvalidFormatError) Error() string {
	return e.text
}

func NewEmailInvalidFormatError() error {
	return &EmailInvalidFormatError{
		text: "email format is invalid",
	}
}

type EmptyFieldError struct {
	text  string
	Field string
}

func (e *EmptyFieldError) Error() string {
	return e.text
}

func NewEmptyFieldError(field string) error {
	text := fmt.Sprintf("%s field is required", strings.ToLower(field))
	return &EmptyFieldError{
		text:  text,
		Field: field,
	}
}
