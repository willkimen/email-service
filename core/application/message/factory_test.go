package message_test

import (
	"emailservice/core/application/apperrors"
	"emailservice/core/application/message"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMessage_Success(t *testing.T) {
	variables := map[string]any{"verification_code": "123456"}

	for _, validType := range validTypes {
		t.Run("Success_With_Type_"+validType, func(t *testing.T) {
			actual_message, actual_err := message.NewMessage(
				id, to, validType, variables,
			)

			assert.Nil(t, actual_err)
			assert.NotNil(t, actual_message)
			assert.Equal(t, id, actual_message.Id)
			assert.Equal(t, to, actual_message.To)
			assert.Equal(t, variables, actual_message.Variables)
			assert.Equal(t, validType, actual_message.Type)
		})
	}
}

func TestNewMessage_RequiredFieldsError(t *testing.T) {
	for _, tc := range requiredFieldCases {
		t.Run(tc.name, func(t *testing.T) {
			actual_message, actual_err := message.NewMessage(
				tc.id, tc.to, tc.messageType, nil,
			)

			assert.Nil(t, actual_message)

			var targetErr *apperrors.InvalidFieldError
			if assert.ErrorAs(t, actual_err, &targetErr) {
				assert.Contains(t, actual_err.Error(), "field is required")
			}
		})
	}
}

func TestNewMessage_InvalidTypeError(t *testing.T) {
	for _, tc := range invalidTypeCases {
		t.Run(tc.name, func(t *testing.T) {
			actual_message, actual_err := message.NewMessage(
				id, to, tc.invalidType, nil,
			)

			assert.Nil(t, actual_message)

			var targetErr *apperrors.InvalidFieldError
			if assert.ErrorAs(t, actual_err, &targetErr) {
				assert.Contains(t, actual_err.Error(), typeFieldName)

			}
		})
	}
}

func TestNewMessage_InvalidEmailError(t *testing.T) {
	for _, tc := range invalidEmailCases {
		t.Run(tc.name, func(t *testing.T) {
			msg, actual_err := message.NewMessage(id, tc.invalidEmail, messageType, nil)

			assert.Nil(t, msg)

			var targetErr *apperrors.InvalidFieldError
			if assert.ErrorAs(t, actual_err, &targetErr) {
				assert.Contains(t, actual_err.Error(), "invalid email format")

			}
		})
	}
}

func TestNewMessage_VerificationCodeRequiredError(t *testing.T) {
	msg, actual_err := message.NewMessage(id, to, messageType, nil)

	assert.Nil(t, msg)

	var targetErr *apperrors.InvalidFieldError
	if assert.ErrorAs(t, actual_err, &targetErr) {
		assert.Contains(
			t,
			actual_err.Error(),
			"cannot be null or empty for this message type",
		)
	}
}
