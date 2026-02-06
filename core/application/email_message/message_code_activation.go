package emailmessage

// ActivationCode represents the email message used to deliver an account
// activation code and related activation details to the user.
type ActivationCode struct {
	BaseMessage
	BaseCodeMessage

	// ActivationLink is the URL the user must access to activate the account.
	ActivationLink string

	// ActivationDeadlineDays defines how many days the activation remains valid.
	ActivationDeadlineDays string
}

// TemplateID returns the identifier of the email template associated
// with the activation code message.
func (ActivationCode) TemplateID() string {
	return TemplateActivationCodeID
}

// GetBodyData returns the data structure used to populate the email template
// associated with the entity.
//
// The returned value contains only the fields required by the template
// renderer and represents a read-only projection of the entity.
// This method does not apply formatting or validation logic;
// it simply exposes the data needed for template interpolation.
func (a *ActivationCode) GetBodyData() any {
	return struct {
		VerificationCode       string
		CodeExpirationHours    string
		ActivationLink         string
		ActivationDeadlineDays string
	}{
		VerificationCode:       a.VerificationCode,
		CodeExpirationHours:    a.CodeExpirationHours,
		ActivationLink:         a.ActivationLink,
		ActivationDeadlineDays: a.ActivationDeadlineDays,
	}
}

func NewActivationCode(
	to, subject, verificationCode, activationLink,
	codeExpirationHours, activationDeadlineDays string,
) *ActivationCode {
	activationCode := &ActivationCode{}

	activationCode.To = to
	activationCode.Subject = subject
	activationCode.VerificationCode = verificationCode
	activationCode.CodeExpirationHours = codeExpirationHours
	activationCode.ActivationLink = activationLink
	activationCode.ActivationDeadlineDays = activationDeadlineDays

	return activationCode
}
