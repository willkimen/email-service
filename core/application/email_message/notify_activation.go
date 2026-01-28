package emailmessage

type NotifyActivation struct {
	Base
	LoginLink string
}

func (NotifyActivation) TemplateID() string {
	return TemplateNotifyActivationID
}

func NewNotifiyActivation(
	to, subject, loginLink string,
) *NotifyActivation {
	notify := &NotifyActivation{}

	notify.To = to
	notify.Subject = subject
	notify.LoginLink = loginLink

	return notify
}
