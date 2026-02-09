package emailmessage

type ResetPasswordCodeBody struct {
	BaseCodeMessage

	// ResetPasswordLink is the URL the user must access to complete
	// the password reset process.
	ResetPasswordLink string
}

// ResetPasswordCode represents the data required to send an email
// containing a verification code and link for password reset.
type ResetPasswordCode struct {
	BaseMessage
	ResetPasswordCodeBody
}

func (ResetPasswordCode) GetEmailType() string {
	return EmailTypeResetPasswordCode
}

func (r *ResetPasswordCode) GetBodyData() any {
	return r.ResetPasswordCodeBody
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
