package outputport

type SendEmailService interface {
	SendEmail(to, subject, body string) error
}
