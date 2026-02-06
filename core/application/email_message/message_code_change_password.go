package emailmessage

// ChangePasswordCode represents the data required to send an email
// containing a verification code for password change operations.
type ChangePasswordCode struct {
	BaseMessage
	BaseCodeMessage
}

// TemplateID returns the identifier of the email template
// associated with password change code messages.
func (ChangePasswordCode) TemplateID() string {
	return TemplateChangePasswordCodeID
}

// GetBodyData returns the data structure used to populate the email template
// associated with the entity.
//
// The returned value contains only the fields required by the template
// renderer and represents a read-only projection of the entity.
// This method does not apply formatting or validation logic;
// it simply exposes the data needed for template interpolation.
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
