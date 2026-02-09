package emailmessage

type ChangeEmailCodeBody struct {
	BaseCodeMessage
}

// ChangeEmailCode represents the data required to send an email
// containing a verification code for changing the user's email address.
type ChangeEmailCode struct {
	BaseMessage
	ChangeEmailCodeBody
}

func (ChangeEmailCode) GetEmailType() string {
	return EmailTypeChangeEmailCode
}

func (c *ChangeEmailCode) GetBodyData() any {
	return c.ChangeEmailCodeBody
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
