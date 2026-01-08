package payload

type ChangeEmailCodePayload struct {
	BasePayload
	BaseCodePayload
}

func NewChangeEmailCodePayload(
	to, subject, verificationCode, templatePath string,
	codeExpirationHours int,
) (*ChangeEmailCodePayload, error) {
	payload := &ChangeEmailCodePayload{}

	payload.To = to
	payload.Subject = subject
	payload.VerificationCode = verificationCode
	payload.CodeExpirationHours = codeExpirationHours

	template, err := RenderHTML(payload, templatePath)
	if err != nil {
		return nil, err
	}

	payload.Body = template
	return payload, nil
}

type NotifyChangeEmailPayload struct {
	BasePayload
	LoginLink string
}

func NewNotifyChangeEmailPayload(
	to, subject, loginLink, templatePath string,
) (*NotifyChangeEmailPayload, error) {
	payload := &NotifyChangeEmailPayload{}

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
