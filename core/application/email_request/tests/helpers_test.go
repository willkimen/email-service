package emailrequest_test

import (
	"emailservice/core/application/email_request"
)

const (
	to                     = "doe@email.com"
	subject                = "some subject"
	verificationCode       = "1234"
	codeExpiratinoHours    = "2"
	link                   = "www.some.com/some"
	activationDeadlineDays = "7"
)

func validActivationCode() *emailrequest.ActivationCode {
	return emailrequest.NewActivationCode(
		to, subject, verificationCode, link,
		codeExpiratinoHours, activationDeadlineDays,
	)
}

func validNotifyActivation() *emailrequest.NotifyActivation {
	return emailrequest.NewNotifiyActivation(
		to, subject, link,
	)
}

func validChangeEmailCode() *emailrequest.ChangeEmailCode {
	return emailrequest.NewChangeEmailCode(
		to, subject, verificationCode, codeExpiratinoHours,
	)
}

func validNotifyChangeEmail() *emailrequest.NotifyChangeEmail {
	return emailrequest.NewNotifyChangeEmail(
		to, subject, link,
	)
}

func validResetPasswordCode() *emailrequest.ResetPasswordCode {
	return emailrequest.NewResetPasswordCode(
		to, subject, verificationCode, link, codeExpiratinoHours,
	)
}

func validNotifyResetPassword() *emailrequest.NotifyResetPassword {
	return emailrequest.NewNotifyResetPassword(
		to, subject, link,
	)
}

func validDeletionCode() *emailrequest.DeletionCode {
	return emailrequest.NewDeletionCode(
		to, subject, verificationCode, codeExpiratinoHours,
	)
}

func validNotifyDeletion() *emailrequest.NotifyDeletion {
	return emailrequest.NewNotifyDeletion(
		to, subject,
	)
}

func validChangePasswordCode() *emailrequest.ChangePasswordCode {
	return emailrequest.NewChangePasswordCode(
		to, subject, verificationCode, codeExpiratinoHours,
	)
}

func validNotifyChangePassword() *emailrequest.NotifyChangePassword {
	return emailrequest.NewNotifyChangePassword(
		to, subject, link,
	)
}
