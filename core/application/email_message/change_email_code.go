package emailmessage

const TemplateChangeEmailCodeID = "change_email_code"

type ChangeEmailCode struct {
	Base
	BaseCode
}

func (ChangeEmailCode) TemplateID() string {
	return TemplateChangeEmailCodeID
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
