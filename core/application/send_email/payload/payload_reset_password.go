package payload

type PasswordResetCodePayload struct {
	BasePayload
	BaseCodePayload
	ResetPasswordLink string
}

func NewPasswordResetCodePayload(
	to, subject, verificationCode, resetPasswordLink, templatePath string,
	codeExpirationHours int,
) (*PasswordResetCodePayload, error) {
	payload := &PasswordResetCodePayload{}

	payload.To = to
	payload.Subject = subject
	payload.VerificationCode = verificationCode
	payload.CodeExpirationHours = codeExpirationHours
	payload.ResetPasswordLink = resetPasswordLink

	template, err := RenderHTML(payload, templatePath)
	if err != nil {
		return nil, err
	}

	payload.Body = template
	return payload, nil
}

type NotifyResetPasswordPayload struct {
	BasePayload
	LoginLink string
}

func NewNotifyResetPasswordPayload(
	to, subject, loginLink, templatePath string,
) (*NotifyResetPasswordPayload, error) {
	payload := &NotifyResetPasswordPayload{}

	payload.To = to
	payload.Subject = subject
	payload.LoginLink = loginLink

	template, err := RenderHTML(payload, templatePath)
	if err != nil {
		return nil, err
	}

	payload.Body = template
	return payload, nil
}
