package payload

type DeletionCodePayload struct {
	BasePayload
	BaseCodePayload
}

func NewDeletionCodePayload(
	to, subject, verificationCode, templatePath string,
	codeExpirationHours int,
) (*DeletionCodePayload, error) {
	payload := &DeletionCodePayload{}

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

type NotifyDeletionPayload struct {
	BasePayload
}

func NewNotifyDeletionPayload(
	to, subject, templatePath string,
) (*NotifyDeletionPayload, error) {
	payload := &NotifyDeletionPayload{}

	payload.To = to
	payload.Subject = subject

	template, err := RenderHTML(payload, templatePath)
	if err != nil {
		return nil, err
	}

	payload.Body = template
	return payload, nil
}
