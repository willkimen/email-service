package emailmessage

// ResetPasswordCode represents the data required to send an email
// containing a verification code and link for password reset.
type ResetPasswordCode struct {
	BaseMessage
	BaseCodeMessage

	// ResetPasswordLink is the URL the user must access to complete
	// the password reset process.
	ResetPasswordLink string
}

// TemplateID returns the identifier of the email template
// associated with password reset code messages.
func (ResetPasswordCode) TemplateID() string {
	return TemplateResetPasswordCodeID
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
