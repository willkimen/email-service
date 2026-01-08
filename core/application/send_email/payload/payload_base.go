package payload

type BasePayload struct {
	To      string
	Subject string
	Body    string
}

type BaseCodePayload struct {
	VerificationCode    string
	CodeExpirationHours int
}
