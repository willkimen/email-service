package emailmessage

type ResetPasswordCode struct {
	Base
	BaseCode
	ResetPasswordLink string
}

func (ResetPasswordCode) TemplateID() string {
	return TemplateResetPasswordCodeID
}

func NewResetPasswordCode(
	to, subject, verificationCode, resetPasswordLink, codeExpirationHours string,
) *ResetPasswordCode {
	notify := &ResetPasswordCode{}

	notify.To = to
	notify.Subject = subject
	notify.VerificationCode = verificationCode
	notify.CodeExpirationHours = codeExpirationHours
	notify.ResetPasswordLink = resetPasswordLink

	return notify
}
