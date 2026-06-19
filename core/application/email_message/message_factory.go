package emailmessage

func NewEmailVerificationCode(
	to, subject, verificationCode string,
) *EmailVerificationCode {
	emailVerificationCode := &EmailVerificationCode{}

	emailVerificationCode.To = to
	emailVerificationCode.Subject = subject
	emailVerificationCode.VerificationCode = verificationCode

	return emailVerificationCode
}

func NewChangeEmailCode(
	to, subject, verificationCode string,
) *ChangeEmailCode {
	changeEmail := &ChangeEmailCode{}

	changeEmail.To = to
	changeEmail.Subject = subject
	changeEmail.VerificationCode = verificationCode

	return changeEmail
}

func NewChangePasswordCode(
	to, subject, verificationCode string,
) *ChangePasswordCode {
	changePassword := &ChangePasswordCode{}

	changePassword.To = to
	changePassword.Subject = subject
	changePassword.VerificationCode = verificationCode

	return changePassword
}

func NewDeletionCode(
	to, subject, verificationCode string,
) *DeletionCode {
	deletion := &DeletionCode{}

	deletion.To = to
	deletion.Subject = subject
	deletion.VerificationCode = verificationCode

	return deletion
}

func NewResetPasswordCode(
	to, subject, verificationCode string,
) *ResetPasswordCode {
	reset := &ResetPasswordCode{}

	reset.To = to
	reset.Subject = subject
	reset.VerificationCode = verificationCode

	return reset
}

func NewNotifyEmailVerification(
	to, subject string,
) *NotifyEmailVerification {
	notify := &NotifyEmailVerification{}

	notify.To = to
	notify.Subject = subject

	return notify
}

func NewNotifyChangeEmail(
	to, subject string,
) *NotifyChangeEmail {
	notify := &NotifyChangeEmail{}

	notify.To = to
	notify.Subject = subject

	return notify
}

func NewNotifyChangePassword(
	to, subject string,
) *NotifyChangePassword {
	notify := &NotifyChangePassword{}

	notify.To = to
	notify.Subject = subject

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
	to, subject string,
) *NotifyResetPassword {
	notify := &NotifyResetPassword{}

	notify.To = to
	notify.Subject = subject

	return notify
}
