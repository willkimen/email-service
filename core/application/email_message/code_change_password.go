package emailmessage

// ChangePasswordCode represents the data required to send an email
// containing a verification code for password change operations.
type ChangePasswordCode struct {
	Base
	BaseCode
}

// TemplateID returns the identifier of the email template
// associated with password change code messages.
func (ChangePasswordCode) TemplateID() string {
	return TemplateChangePasswordCodeID
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

