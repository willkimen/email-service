package emailmessage

type ChangePasswordCodeBody struct {
	BaseCodeMessage
}

// ChangePasswordCode represents the data required to send an email
// containing a verification code for password change operations.
type ChangePasswordCode struct {
	BaseMessage
	ChangeEmailCodeBody
}

func (ChangePasswordCode) GetEmailType() string {
	return EmailTypeChangePasswordCode
}

func (c *ChangePasswordCode) GetBodyData() any {
	return c.ChangeEmailCodeBody
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
