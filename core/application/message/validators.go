package message

import (
	"emailservice/core/application/apperrors"
	"fmt"
	"regexp"
	"strings"
)

const (
	IdFieldName             = "id"
	ToFieldName             = "to"
	TypeFieldName           = "type"
	VariablesFieldName      = "variables"
	VerificationCodeKeyName = "verification_code"
)

// FieldRule binds a raw value to its logical field name,
// allowing validation errors to report which field is invalid.
type FieldRule struct {
	Value     string
	FieldName string
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// ValidateEmailFormat ensures the email address follows a valid format.
// Only syntactically valid email addresses are accepted.
//
// Errors:
//   - apperrors.InvalidFieldError: if the email syntax is invalid.
func ValidateEmailFormat(email, fieldName string) error {
	if !emailRegex.MatchString(email) {
		return apperrors.NewEmailInvalidFormatError(fieldName)
	}
	return nil
}

// ValidateRequiredFields ensures that all provided fields contain data.
// Empty or whitespace-only values are not accepted.
//
// Errors:
//   - apperrors.InvalidFieldError: for the first field encountered that is
//     empty or contains only whitespace.
func ValidateRequiredFields(fields ...FieldRule) error {
	for _, f := range fields {
		if strings.TrimSpace(f.Value) == "" {
			return apperrors.NewEmptyFieldError(f.FieldName)
		}
	}
	return nil
}

// ValidateType ensures the message type is valid and supported by the system.
// Unsupported or unrecognized message types are not accepted.
//
// Errors:
//   - apperrors.InvalidFieldError: if the message type is not supported.
func ValidateType(messageType, fieldName string) error {
	switch messageType {
	case MessageTypeEmailVerificationCode, MessageTypeNotifyEmailVerified,
		MessageTypeChangeEmailCode, MessageTypeNotifyEmailChanged,
		MessageTypeChangePasswordCode, MessageTypeNotifyPasswordChanged,
		MessageTypeResetPasswordCode, MessageTypeNotifyPasswordReset,
		MessageTypeAccountDeletionCode, MessageTypeNotifyAccountDeleted:
		return nil
	default:
		return apperrors.NewInvalidTypeError(fieldName)
	}
}

// ValidateVerificationCodeVariables ensures that the required verification code
// exists and is valid within the variables map, depending on the current message type.
//
// message types that do not require a verification code will skip this validation.
//
// Errors:
//   - apperrors.InvalidFieldError: if variables is null or map empty, missing the verification_code,
//     or if the verification_code is not a non-empty string.
func ValidateVerificationCodeVariables(messageType string, variables map[string]any) error {
	switch messageType {
	case MessageTypeEmailVerificationCode, MessageTypeChangeEmailCode,
		MessageTypeChangePasswordCode, MessageTypeResetPasswordCode,
		MessageTypeAccountDeletionCode:

		mapEmpty := 0
		if variables == nil || len(variables) == mapEmpty {
			return apperrors.NewVerificationCodeMissingError(
				VariablesFieldName,
				fmt.Sprintf(
					"%s cannot be null or empty for this message type",
					VariablesFieldName,
				),
			)
		}

		rawCode, exists := variables[VerificationCodeKeyName]
		if !exists {
			return apperrors.NewVerificationCodeMissingError(
				VerificationCodeKeyName,
				fmt.Sprintf("%s is required", VerificationCodeKeyName),
			)
		}

		code, ok := rawCode.(string)
		if !ok || strings.TrimSpace(code) == "" {
			return apperrors.NewVerificationCodeMissingError(
				VerificationCodeKeyName,
				fmt.Sprintf(
					"%s must be a non-empty string",
					VerificationCodeKeyName),
			)
		}
	}

	return nil
}
