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

func (ActivationCode) GetEmailType() string {
	return EmailTypeActivationCode
}

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
