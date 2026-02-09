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

func (ResetPasswordCode) GetEmailType() string {
	return EmailTypeResetPasswordCode
}

func (r *ResetPasswordCode) GetBodyData() any {
	return struct {
		VerificationCode    string
		CodeExpirationHours string
		ResetPasswordLink   string
	}{
		VerificationCode:    r.VerificationCode,
		CodeExpirationHours: r.CodeExpirationHours,
		ResetPasswordLink:   r.ResetPasswordLink,
	}
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
