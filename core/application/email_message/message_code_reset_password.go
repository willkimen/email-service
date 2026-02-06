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

// GetBodyData returns the data structure used to populate the email template
// associated with the entity.
//
// The returned value contains only the fields required by the template
// renderer and represents a read-only projection of the entity.
// This method does not apply formatting or validation logic;
// it simply exposes the data needed for template interpolation.
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
