package emailmessage

type EmailMessage interface {
	ValidateData() error
}

type Base struct {
	To      string
	Subject string
}

type BaseCode struct {
	VerificationCode    string
	CodeExpirationHours string
}
