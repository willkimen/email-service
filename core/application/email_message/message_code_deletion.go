package emailmessage

// DeletionCode represents the data required to send an email
// containing a verification code for account deletion operations.
type DeletionCode struct {
	BaseMessage
	BaseCodeMessage
}

// TemplateID returns the identifier of the email template
// associated with account deletion code messages.
func (DeletionCode) TemplateID() string {
	return TemplateDeletionCodeID
}

// GetBodyData returns the data structure used to populate the email template
// associated with the entity.
//
// The returned value contains only the fields required by the template
// renderer and represents a read-only projection of the entity.
// This method does not apply formatting or validation logic;
// it simply exposes the data needed for template interpolation.
func (d *DeletionCode) GetBodyData() any {
	return struct {
		VerificationCode    string
		CodeExpirationHours string
	}{
		VerificationCode:    d.VerificationCode,
		CodeExpirationHours: d.CodeExpirationHours,
	}
}

func NewDeletionCode(
	to, subject, verificationCode, codeExpirationHours string,
) *DeletionCode {
	deletion := &DeletionCode{}

	deletion.To = to
	deletion.Subject = subject
	deletion.VerificationCode = verificationCode
	deletion.CodeExpirationHours = codeExpirationHours

	return deletion
}
