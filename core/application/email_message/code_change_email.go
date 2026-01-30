package emailmessage

// ChangeEmailCode represents the data required to send an email
// containing a verification code for changing the user's email address.
type ChangeEmailCode struct {
	Base
	BaseCode
}

// TemplateID returns the identifier of the email template
// used to render the change email code message.
func (ChangeEmailCode) TemplateID() string {
	return TemplateChangeEmailCodeID
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

