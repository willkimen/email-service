package emailmessage

// ChangePasswordCode represents the data required to send an email
// containing a verification code for password change operations.
type ChangePasswordCode struct {
	BaseMessage
	BaseCodeMessage
}

func (ChangePasswordCode) GetEmailType() string {
	return EmailTypeChangePasswordCode
}

func (c *ChangePasswordCode) GetBodyData() any {
	return struct {
		VerificationCode    string
		CodeExpirationHours string
	}{
		VerificationCode:    c.VerificationCode,
		CodeExpirationHours: c.CodeExpirationHours,
	}
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
