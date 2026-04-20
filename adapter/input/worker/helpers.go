package worker

import (
	"emailservice/core/application/email_message"
	"encoding/json"
	"errors"
)

// ToEmailMessage converts a serialized task payload into a concrete
// EmailMessage  object.
//
// This function acts as an infrastructure-level mapper responsible for
// reconstructing the appropriate email message type based on the
// EmailType discriminator contained in the payload.
//
// The payload is expected to contain:
// - Basic metadata (To, Subject, EmailType)
// - A BodyData field encoded as raw JSON
//
// For each supported EmailType, the function unmarshals the BodyData
// into the corresponding body structure and invokes the appropriate
// factory constructor.
//
// If the payload cannot be deserialized, or if the EmailType is unknown,
// an error is returned. The caller is responsible for deciding whether
// the error should trigger retries or be treated as a permanent failure.
func ToEmailMessage(payloadBytes []byte) (emailmessage.EmailMessage, error) {
	var raw struct {
		To        string
		Subject   string
		EmailType string
		BodyData  json.RawMessage
	}

	if err := json.Unmarshal(payloadBytes, &raw); err != nil {
		return nil, err
	}

	switch raw.EmailType {

	case emailmessage.EmailTypeEmailVerificationCode:
		var body emailmessage.EmailVerificationCodeBody
		if err := json.Unmarshal(raw.BodyData, &body); err != nil {
			return nil, err
		}

		return emailmessage.NewEmailVerificationCode(
			raw.To,
			raw.Subject,
			body.VerificationCode,
			body.EmailVerificationLink,
			body.CodeExpirationHours,
			body.EmailVerificationDeadlineDays,
		), nil

	case emailmessage.EmailTypeNotifyEmailVerification:
		var body emailmessage.NotifyEmailVerificationBody
		if err := json.Unmarshal(raw.BodyData, &body); err != nil {
			return nil, err
		}

		return emailmessage.NewNotifyEmailVerification(
			raw.To,
			raw.Subject,
			body.LoginLink,
		), nil

	case emailmessage.EmailTypeChangeEmailCode:
		var body emailmessage.ChangeEmailCodeBody
		if err := json.Unmarshal(raw.BodyData, &body); err != nil {
			return nil, err
		}

		return emailmessage.NewChangeEmailCode(
			raw.To,
			raw.Subject,
			body.VerificationCode,
			body.CodeExpirationHours,
		), nil

	case emailmessage.EmailTypeNotifyChangeEmail:
		var body emailmessage.NotifyChangeEmailBody
		if err := json.Unmarshal(raw.BodyData, &body); err != nil {
			return nil, err
		}

		return emailmessage.NewNotifyChangeEmail(
			raw.To,
			raw.Subject,
			body.LoginLink,
		), nil

	case emailmessage.EmailTypeChangePasswordCode:
		var body emailmessage.ChangePasswordCodeBody
		if err := json.Unmarshal(raw.BodyData, &body); err != nil {
			return nil, err
		}

		return emailmessage.NewChangePasswordCode(
			raw.To,
			raw.Subject,
			body.VerificationCode,
			body.CodeExpirationHours,
		), nil

	case emailmessage.EmailTypeNotifyChangePassword:
		var body emailmessage.NotifyChangePasswordBody
		if err := json.Unmarshal(raw.BodyData, &body); err != nil {
			return nil, err
		}

		return emailmessage.NewNotifyChangePassword(
			raw.To,
			raw.Subject,
			body.LoginLink,
		), nil

	case emailmessage.EmailTypeResetPasswordCode:
		var body emailmessage.ResetPasswordCodeBody
		if err := json.Unmarshal(raw.BodyData, &body); err != nil {
			return nil, err
		}

		return emailmessage.NewResetPasswordCode(
			raw.To,
			raw.Subject,
			body.VerificationCode,
			body.ResetPasswordLink,
			body.CodeExpirationHours,
		), nil

	case emailmessage.EmailTypeNotifyResetPassword:
		var body emailmessage.NotifyResetPasswordBody
		if err := json.Unmarshal(raw.BodyData, &body); err != nil {
			return nil, err
		}

		return emailmessage.NewNotifyResetPassword(
			raw.To,
			raw.Subject,
			body.LoginLink,
		), nil

	case emailmessage.EmailTypeDeletionCode:
		var body emailmessage.DeletionCodeBody
		if err := json.Unmarshal(raw.BodyData, &body); err != nil {
			return nil, err
		}

		return emailmessage.NewDeletionCode(
			raw.To,
			raw.Subject,
			body.VerificationCode,
			body.CodeExpirationHours,
		), nil

	case emailmessage.EmailTypeNotifyDeletion:
		var body emailmessage.NotifyDeletionBody
		if err := json.Unmarshal(raw.BodyData, &body); err != nil {
			return nil, err
		}

		return emailmessage.NewNotifyDeletion(
			raw.To,
			raw.Subject,
		), nil
	}

	return nil, errors.New("unknown email type")
}
