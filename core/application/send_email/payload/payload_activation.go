package payload

type ActivationCodePayload struct {
	BasePayload
	BaseCodePayload
	ActivationLink         string
	ActivationDeadlineDays int
}

type NotifyActivationPayload struct {
	BasePayload
	LoginLink string
}

func NewActivationCodePayload(
	to, subject, verificationCode, activationLink string,
	codeExpirationHours, activationDeadlineDays int,
	templatePath string,
) (*ActivationCodePayload, error) {
	payload := &ActivationCodePayload{}

	payload.To = to
	payload.Subject = subject
	payload.VerificationCode = verificationCode
	payload.ActivationLink = activationLink
	payload.CodeExpirationHours = codeExpirationHours
	payload.ActivationDeadlineDays = activationDeadlineDays

	template, err := RenderHTML(payload, templatePath)
	if err != nil {
		return nil, err
	}

	payload.Body = template
	return payload, nil
}

func NewNotifiyActivationPayload(
	to, subject, loginLink, templatePath string,
) (*NotifyActivationPayload, error) {
	payload := &NotifyActivationPayload{}

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
