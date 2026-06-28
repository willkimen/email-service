package message_test

import (
	"emailservice/core/application/apperrors"
	"emailservice/core/application/message"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateEmailFormat_Success(t *testing.T) {
	actual_err := message.ValidateEmailFormat("correct@email.com", toFieldName)
	assert.Nil(t, actual_err)
}

func TestValidateEmailFormat_InvalidEmailError(t *testing.T) {
	for _, tc := range invalidEmailCases {
		t.Run(tc.name, func(t *testing.T) {
			actual_err := message.ValidateEmailFormat(tc.invalidEmail, toFieldName)

			var targetErr *apperrors.InvalidFieldError
			if assert.ErrorAs(t, actual_err, &targetErr) {
				assert.Contains(
					t,
					actual_err.Error(),
					"to contains an invalid email format",
				)
			}
		})
	}
}

func TestValidateRequiredFields_Success(t *testing.T) {
	actual_err := message.ValidateRequiredFields(message.FieldRule{
		Value:     "fake-value",
		FieldName: "fake_field",
	})

	assert.Nil(t, actual_err)
}

func TestValidateRequiredFields_EmptyFieldError(t *testing.T) {
	actual_err := message.ValidateRequiredFields(message.FieldRule{
		Value:     "",
		FieldName: "fake_name",
	})

	var targetErr *apperrors.InvalidFieldError
	if assert.ErrorAs(t, actual_err, &targetErr) {
		assert.Contains(t, actual_err.Error(), "fake_name field is required")
	}
}

func TestValidateType_Success(t *testing.T) {
	for _, valid := range validTypes {
		t.Run("ValidType_"+valid, func(t *testing.T) {
			actual_err := message.ValidateType(valid, typeFieldName)
			assert.Nil(t, actual_err)
		})
	}
}

func TestValidateType_InvalidTypeError(t *testing.T) {
	for _, tc := range invalidTypeCases {
		t.Run(tc.name, func(t *testing.T) {
			actual_err := message.ValidateType(tc.invalidType, typeFieldName)

			var targetErr *apperrors.InvalidFieldError
			if assert.ErrorAs(t, actual_err, &targetErr) {
				assert.Contains(t, actual_err.Error(), typeFieldName)

			}
		})
	}
}

func TestValidateVerificationCodeVariables_CodeType_Success(t *testing.T) {
	for _, c := range codeTypes {
		variables := map[string]any{
			"verification_code": "123456",
		}

		t.Run("CodeType_"+c, func(t *testing.T) {
			actual_err := message.ValidateVerificationCodeVariables(
				c,
				variables,
			)
			assert.Nil(t, actual_err)
		})
	}
}

func TestValidateVerificationCodeVariables_IncorrectVariables_Error(t *testing.T) {
	for _, v := range variablesError {
		t.Run(v.name, func(t *testing.T) {
			actual_err := message.ValidateVerificationCodeVariables(
				messageType,
				v.variables,
			)
			require.Error(t, actual_err)
			var fieldError *apperrors.InvalidFieldError
			if assert.ErrorAs(t, actual_err, &fieldError) {
				assert.Contains(t, actual_err.Error(), v.partialMessageError)

			}
		})
	}
}

func TestValidateVerificationCodeVariables_NotifyType_NoError(t *testing.T) {
	for _, c := range notifyTypes {
		variables := map[string]any{
			"verification_code": "123456",
		}

		t.Run("CodeType_"+c, func(t *testing.T) {
			actual_err := message.ValidateVerificationCodeVariables(
				c,
				variables,
			)
			assert.Nil(t, actual_err)
		})
	}
}
