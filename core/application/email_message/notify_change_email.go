package emailmessage

type NotifyChangeEmail struct {
	Base
	LoginLink string
}

func (NotifyChangeEmail) TemplateID() string {
	return TemplateNotifyChangeEmailID
}

func NewNotifyChangeEmail(
	to, subject, loginLink string,
) *NotifyChangeEmail {
	notify := &NotifyChangeEmail{}

	notify.To = to
	notify.Subject = subject
	notify.LoginLink = loginLink

	return notify
}
