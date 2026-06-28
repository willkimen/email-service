package apperrors

import (
	"fmt"
	"strings"
)

// NewInfrastructureError creates a new InfrastructureError to wrap
// technical or operational failures from the infrastructure layer.
func NewInfrastructureError(message string, originalCause error) error {
	return &InfrastructureError{
		Message:       message,
		OriginalCause: originalCause,
	}
}

// NewEmailInvalidFormatError creates and initializes a new InvalidFieldError,
// that represents a validation error indicating
// that an email address does not match the expected format.
func NewEmailInvalidFormatError(fieldName string) error {
	message := fmt.Sprintf("%s contains an invalid email format", strings.ToLower(fieldName))
	return &InvalidFieldError{
		Message:   message,
		FieldName: fieldName,
	}
}

// NewEmptyFieldError creates and initializes a new InvalidFieldError,
// that represents a validation error indicating that
// a required field was not provided.
func NewEmptyFieldError(fieldName string) error {
	message := fmt.Sprintf("%s field is required", strings.ToLower(fieldName))
	return &InvalidFieldError{
		Message:   message,
		FieldName: fieldName,
	}
}

// NewInvalidTypeError creates and initializes a new InvalidFieldError,
// that represents a validation error indicating that
// an type field is invalid.
func NewInvalidTypeError(fieldName string) error {
	fieldName = strings.ToLower(fieldName)
	message := fmt.Sprintf("%s is not a valid message type", fieldName)

	return &InvalidFieldError{
		Message:   message,
		FieldName: fieldName,
	}
}

// NewVerificationCodeMissingError creates and initializes a new InvalidFieldError,
// that represents a validation error indicating that the verification code variable
// is missing, null, or invalid for the current message type.
func NewVerificationCodeMissingError(fieldName, message string) error {
	return &InvalidFieldError{
		Message:   message,
		FieldName: strings.ToLower(fieldName),
	}
}
