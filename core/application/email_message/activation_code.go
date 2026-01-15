package emailmessage

const TemplateActivationCodeID = "activation_code"

type ActivationCode struct {
	Base
	BaseCode
	ActivationLink         string
	ActivationDeadlineDays string
	TemplateId             string
}

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
	activationCode.ActivationLink = activationLink
	activationCode.CodeExpirationHours = codeExpirationHours
	activationCode.ActivationDeadlineDays = activationDeadlineDays

	return activationCode
}
