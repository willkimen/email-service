package input

import (
	"emailservice/core/application/send_email/payload"
)

type SendEmailUserCase interface {
	SendActivationCode(input *payload.ActivationCodePayload) error
	NotifyActivation(input *payload.NotifyActivationPayload) error

	SendChangeEmailCode(input *payload.ChangeEmailCodePayload) error
	NotifyChangeEmail(input *payload.NotifyChangeEmailPayload) error

	SendResetPasswordCode(input *payload.PasswordResetCodePayload) error
	NotifyResetPassword(input *payload.NotifyResetPasswordPayload) error

	SendDeletionCode(input *payload.DeletionCodePayload) error
	NotifyDeletion(input *payload.NotifyDeletionPayload) error
}
