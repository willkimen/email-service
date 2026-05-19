package emailmessage_test

import (
	"emailservice/core/application/email_message"
)

const (
	to                     = "doe@email.com"
	subject                = "some subject"
	verificationCode       = "1234"
	codeExpirationTime    = "2"
	emailVerificationDeadlineDays = "7"
)

func validEmailVerificationCode() *emailmessage.EmailVerificationCode {
	return emailmessage.NewEmailVerificationCode(
		to, subject, verificationCode,
		codeExpirationTime, emailVerificationDeadlineDays,
	)
}

func validNotifyEmailVerification() *emailmessage.NotifyEmailVerification {
	return emailmessage.NewNotifyEmailVerification(
		to, subject,
	)
}

func validChangeEmailCode() *emailmessage.ChangeEmailCode {
	return emailmessage.NewChangeEmailCode(
		to, subject, verificationCode, codeExpirationTime,
	)
}

func validNotifyChangeEmail() *emailmessage.NotifyChangeEmail {
	return emailmessage.NewNotifyChangeEmail(
		to, subject,
	)
}

func validResetPasswordCode() *emailmessage.ResetPasswordCode {
	return emailmessage.NewResetPasswordCode(
		to, subject, verificationCode, codeExpirationTime,
	)
}

func validNotifyResetPassword() *emailmessage.NotifyResetPassword {
	return emailmessage.NewNotifyResetPassword(
		to, subject,
	)
}

func validDeletionCode() *emailmessage.DeletionCode {
	return emailmessage.NewDeletionCode(
		to, subject, verificationCode, codeExpirationTime,
	)
}

func validNotifyDeletion() *emailmessage.NotifyDeletion {
	return emailmessage.NewNotifyDeletion(
		to, subject,
	)
}

func validChangePasswordCode() *emailmessage.ChangePasswordCode {
	return emailmessage.NewChangePasswordCode(
		to, subject, verificationCode, codeExpirationTime,
	)
}

func validNotifyChangePassword() *emailmessage.NotifyChangePassword {
	return emailmessage.NewNotifyChangePassword(
		to, subject,
	)
}
