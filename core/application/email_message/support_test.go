package emailmessage_test

import (
	"emailservice/core/application/email_message"
)

const (
	to                     = "doe@email.com"
	subject                = "some subject"
	verificationCode       = "1234"
	codeExpiratinoHours    = "2"
	link                   = "www.some.com/some"
	emailVerificationDeadlineDays = "7"
)

func validEmailVerificationCode() *emailmessage.EmailVerificationCode {
	return emailmessage.NewEmailVerificationCode(
		to, subject, verificationCode, link,
		codeExpiratinoHours, emailVerificationDeadlineDays,
	)
}

func validNotifyEmailVerification() *emailmessage.NotifyEmailVerification {
	return emailmessage.NewNotifyEmailVerification(
		to, subject, link,
	)
}

func validChangeEmailCode() *emailmessage.ChangeEmailCode {
	return emailmessage.NewChangeEmailCode(
		to, subject, verificationCode, codeExpiratinoHours,
	)
}

func validNotifyChangeEmail() *emailmessage.NotifyChangeEmail {
	return emailmessage.NewNotifyChangeEmail(
		to, subject, link,
	)
}

func validResetPasswordCode() *emailmessage.ResetPasswordCode {
	return emailmessage.NewResetPasswordCode(
		to, subject, verificationCode, link, codeExpiratinoHours,
	)
}

func validNotifyResetPassword() *emailmessage.NotifyResetPassword {
	return emailmessage.NewNotifyResetPassword(
		to, subject, link,
	)
}

func validDeletionCode() *emailmessage.DeletionCode {
	return emailmessage.NewDeletionCode(
		to, subject, verificationCode, codeExpiratinoHours,
	)
}

func validNotifyDeletion() *emailmessage.NotifyDeletion {
	return emailmessage.NewNotifyDeletion(
		to, subject,
	)
}

func validChangePasswordCode() *emailmessage.ChangePasswordCode {
	return emailmessage.NewChangePasswordCode(
		to, subject, verificationCode, codeExpiratinoHours,
	)
}

func validNotifyChangePassword() *emailmessage.NotifyChangePassword {
	return emailmessage.NewNotifyChangePassword(
		to, subject, link,
	)
}
