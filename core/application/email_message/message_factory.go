package emailmessage

func NewEmailVerificationCode(
	to, subject, verificationCode, emailVerificationLink,
	codeExpirationTime, emailVerificationDeadlineDays string,
) *EmailVerificationCode {
	emailVerificationCode := &EmailVerificationCode{}

	emailVerificationCode.To = to
	emailVerificationCode.Subject = subject
	emailVerificationCode.VerificationCode = verificationCode
	emailVerificationCode.CodeExpirationTime = codeExpirationTime
	emailVerificationCode.EmailVerificationLink = emailVerificationLink
	emailVerificationCode.EmailVerificationDeadlineDays = emailVerificationDeadlineDays

	return emailVerificationCode
}

func NewChangeEmailCode(
	to, subject, verificationCode, codeExpirationTime string,
) *ChangeEmailCode {
	changeEmail := &ChangeEmailCode{}

	changeEmail.To = to
	changeEmail.Subject = subject
	changeEmail.VerificationCode = verificationCode
	changeEmail.CodeExpirationTime = codeExpirationTime

	return changeEmail
}

func NewChangePasswordCode(
	to, subject, verificationCode, codeExpirationTime string,
) *ChangePasswordCode {
	changePassword := &ChangePasswordCode{}

	changePassword.To = to
	changePassword.Subject = subject
	changePassword.VerificationCode = verificationCode
	changePassword.CodeExpirationTime = codeExpirationTime

	return changePassword
}

func NewDeletionCode(
	to, subject, verificationCode, codeExpirationTime string,
) *DeletionCode {
	deletion := &DeletionCode{}

	deletion.To = to
	deletion.Subject = subject
	deletion.VerificationCode = verificationCode
	deletion.CodeExpirationTime = codeExpirationTime

	return deletion
}

func NewResetPasswordCode(
	to, subject, verificationCode, resetPasswordLink, codeExpirationTime string,
) *ResetPasswordCode {
	reset := &ResetPasswordCode{}

	reset.To = to
	reset.Subject = subject
	reset.VerificationCode = verificationCode
	reset.CodeExpirationTime = codeExpirationTime
	reset.ResetPasswordLink = resetPasswordLink

	return reset
}

func NewNotifyEmailVerification(
	to, subject, loginLink string,
) *NotifyEmailVerification {
	notify := &NotifyEmailVerification{}

	notify.To = to
	notify.Subject = subject
	notify.LoginLink = loginLink

	return notify
}

func NewNotifyChangeEmail(
	to, subject, loginLink string,
) *NotifyChangeEmail {
	notify := &NotifyChangeEmail{}

	notify.To = to
	notify.Subject = subject
	notify.LoginLink = loginLink

	return notify
}

func NewNotifyChangePassword(
	to, subject, loginLink string,
) *NotifyChangePassword {
	notify := &NotifyChangePassword{}

	notify.To = to
	notify.Subject = subject
	notify.LoginLink = loginLink

	return notify
}

func NewNotifyDeletion(
	to, subject string,
) *NotifyDeletion {
	notify := &NotifyDeletion{}

	notify.To = to
	notify.Subject = subject

	return notify
}

func NewNotifyResetPassword(
	to, subject, loginLink string,
) *NotifyResetPassword {
	notify := &NotifyResetPassword{}

	notify.To = to
	notify.Subject = subject
	notify.LoginLink = loginLink

	return notify
}
