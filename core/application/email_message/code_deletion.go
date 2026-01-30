package emailmessage

// DeletionCode represents the data required to send an email
// containing a verification code for account deletion operations.
type DeletionCode struct {
	Base
	BaseCode
}

// TemplateID returns the identifier of the email template
// associated with account deletion code messages.
func (DeletionCode) TemplateID() string {
	return TemplateDeletionCodeID
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

