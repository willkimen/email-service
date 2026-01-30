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
