package emailmessage

const TemplateChangePasswordCodeID = "change_password_code"

type ChangePasswordCode struct {
	Base
	BaseCode
	TemplateId string
}

func (ChangePasswordCode) TemplateID() string {
	return TemplateChangePasswordCodeID
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
