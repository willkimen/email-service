package emailmessage

func NewEmailVerificationCode(
	to, subject, verificationCode, emailVerificationLink,
	codeExpirationHours, emailVerificationDeadlineDays string,
) *EmailVerificationCode {
	emailVerificationCode := &EmailVerificationCode{}

	emailVerificationCode.To = to
	emailVerificationCode.Subject = subject
	emailVerificationCode.VerificationCode = verificationCode
	emailVerificationCode.CodeExpirationHours = codeExpirationHours
	emailVerificationCode.EmailVerificationLink = emailVerificationLink
	emailVerificationCode.EmailVerificationDeadlineDays = emailVerificationDeadlineDays

	return emailVerificationCode
}

func NewChangeEmailCode(
	to, subject, verificationCode, codeExpirationHours string,
) *ChangeEmailCode {
	changeEmail := &ChangeEmailCode{}

	changeEmail.To = to
	changeEmail.Subject = subject
	changeEmail.VerificationCode = verificationCode
	changeEmail.CodeExpirationHours = codeExpirationHours

	return changeEmail
}

func NewChangePasswordCode(
	to, subject, verificationCode, codeExpirationHours string,
) *ChangePasswordCode {
	changePassword := &ChangePasswordCode{}

	changePassword.To = to
	changePassword.Subject = subject
	changePassword.VerificationCode = verificationCode
	changePassword.CodeExpirationHours = codeExpirationHours

	return changePassword
}

func NewDeletionCode(
	to, subject, verificationCode, codeExpirationHours string,
) *DeletionCode {
	deletion := &DeletionCode{}

	deletion.To = to
	deletion.Subject = subject
	deletion.VerificationCode = verificationCode
	deletion.CodeExpirationHours = codeExpirationHours

	return deletion
}

func NewResetPasswordCode(
	to, subject, verificationCode, resetPasswordLink, codeExpirationHours string,
) *ResetPasswordCode {
	reset := &ResetPasswordCode{}

	reset.To = to
	reset.Subject = subject
	reset.VerificationCode = verificationCode
	reset.CodeExpirationHours = codeExpirationHours
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
