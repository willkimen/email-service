package message_test

import (
	"emailservice/core/application/message"
)

const (
	id          = "fake-id"
	to          = "email@email.com"
	messageType = message.MessageTypeAccountDeletionCode
)

const (
	idFieldName             = "id"
	toFieldName             = "to"
	typeFieldName           = "type"
	variablesFieldName      = "variables"
	verificationCodeKeyName = "verification_code"
)

var requiredFieldCases = []struct {
	name        string
	id          string
	to          string
	messageType string
}{
	{
		name:        "empty id",
		id:          "",
		to:          "email@email.com",
		messageType: message.MessageTypeAccountDeletionCode,
	},
	{
		name:        "empty to (email)",
		id:          "id-fake",
		to:          "",
		messageType: message.MessageTypeAccountDeletionCode,
	},
	{
		name:        "empty verification code",
		id:          "id-fake",
		to:          "email@email.com",
		messageType: "",
	},
}

var invalidEmailCases = []struct {
	name         string
	invalidEmail string
}{
	{name: "without @", invalidEmail: "testegmail.com"},
	{name: "without domain", invalidEmail: "teste@"},
	{name: "without user", invalidEmail: "@gmail.com"},
	{name: "without TLD (.com, .org)", invalidEmail: "teste@gmail"},
	{name: "TLD incomplete", invalidEmail: "teste@gmail."},
	{name: "TLD short (less than 2)", invalidEmail: "teste@gmail.c"},
	{name: "with space", invalidEmail: "teste ovelha@gmail.com"},
	{name: "invalid character", invalidEmail: "teste#v@gmail.com"},
	{name: "two @", invalidEmail: "teste@gm@ail.com"},
}

var validTypes = []string{
	message.MessageTypeEmailVerificationCode,
	message.MessageTypeNotifyEmailVerified,
	message.MessageTypeChangeEmailCode,
	message.MessageTypeNotifyEmailChanged,
	message.MessageTypeChangePasswordCode,
	message.MessageTypeNotifyPasswordChanged,
	message.MessageTypeResetPasswordCode,
	message.MessageTypeNotifyPasswordReset,
	message.MessageTypeAccountDeletionCode,
	message.MessageTypeNotifyAccountDeleted,
}

var codeTypes = []string{
	message.MessageTypeEmailVerificationCode,
	message.MessageTypeChangeEmailCode,
	message.MessageTypeChangePasswordCode,
	message.MessageTypeResetPasswordCode,
	message.MessageTypeAccountDeletionCode,
}

var notifyTypes = []string{
	message.MessageTypeNotifyEmailVerified,
	message.MessageTypeNotifyEmailChanged,
	message.MessageTypeNotifyPasswordChanged,
	message.MessageTypeNotifyPasswordReset,
	message.MessageTypeNotifyAccountDeleted,
}

var variablesError = []struct {
	name                string
	variables           map[string]any
	partialMessageError string
}{
	{
		name:                "Nil map",
		variables:           nil,
		partialMessageError: "cannot be null or empty for this message type",
	},
	{
		name:                "Empty map",
		variables:           map[string]any{},
		partialMessageError: "cannot be null or empty for this message type",
	},
	{
		name:                "Ket not found",
		variables:           map[string]any{"other": "value"},
		partialMessageError: "is required",
	},

	{
		name:                "Empty string",
		variables:           map[string]any{"verification_code": ""},
		partialMessageError: "must be a non-empty string",
	},
}

var invalidTypeCases = []struct {
	name        string
	invalidType string
}{
	{name: "Empty type", invalidType: ""},
	{name: "Unknown string", invalidType: "INVALID_EMAIL_TYPE"},
}
