package message

import "strings"

// NewMessage creates and validates a new Message
//
// Returns apperrors.InvalidFieldError:
//   - If any required argument is empty.
//   - If the recipient's email address ("to") is malformed.
//   - If the type field is invalid (non exists).
//   - If the required variables for the specific message type are missing or invalid.
func NewMessage(id, to, messageType string, variables map[string]any) (*Message, error) {
	to = strings.ToLower(to)
	messageType = strings.ToLower(messageType)

	if err := ValidateRequiredFields(
		FieldRule{id, IdFieldName},
		FieldRule{to, ToFieldName},
		FieldRule{messageType, TypeFieldName},
	); err != nil {
		return nil, err
	}

	if err := ValidateType(messageType, TypeFieldName); err != nil {
		return nil, err
	}

	if err := ValidateEmailFormat(to, ToFieldName); err != nil {
		return nil, err
	}

	if err := ValidateVerificationCodeVariables(messageType, variables); err != nil {
		return nil, err
	}

	return &Message{
		Id:        id,
		Type:      messageType,
		To:        to,
		Variables: variables,
	}, nil
}
