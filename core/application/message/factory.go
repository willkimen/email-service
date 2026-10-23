package message

import "strings"

// NewMessage creates and validates a new Message
//
// It is strongly recommended to use this factory function (or workflow-specific
// factories) instead of instantiating the Message struct directly, as it centralizes
// and encapsulates all business validation rules required for a Message entity.
//
// Any provided message type must strictly match one of the system's pre-existing
// constants (e.g., MessageTypeEmailVerificationCode).
//
// Example:
//
//	variables := map[string]any{"code": "123456"}
//	msg, err := message.NewMessage(
//		"msg-123",
//		"user@example.com",
//		message.MessageTypeEmailVerificationCode,
//		variables
//	)
//	if err != nil {
//	    // handle validation error
//	}
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
