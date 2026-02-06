package emailmessage

// ChangeEmailCode represents the data required to send an email
// containing a verification code for changing the user's email address.
type ChangeEmailCode struct {
	BaseMessage
	BaseCodeMessage
}

// TemplateID returns the identifier of the email template
// used to render the change email code message.
func (ChangeEmailCode) TemplateID() string {
	return TemplateChangeEmailCodeID
}

// GetBodyData returns the data structure used to populate the email template
// associated with the entity.
//
// The returned value contains only the fields required by the template
// renderer and represents a read-only projection of the entity.
// This method does not apply formatting or validation logic;
// it simply exposes the data needed for template interpolation.
func (c *ChangeEmailCode) GetBodyData() any {
	return struct {
		VerificationCode    string
		CodeExpirationHours string
	}{
		VerificationCode:    c.VerificationCode,
		CodeExpirationHours: c.CodeExpirationHours,
	}
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
