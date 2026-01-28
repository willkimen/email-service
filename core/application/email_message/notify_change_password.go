package emailmessage

type NotifyChangePassword struct {
	Base
	LoginLink string
}

func (NotifyChangePassword) TemplateID() string {
	return TemplateNotifyChangePasswordID
}

func NewNotifyChangePassword(
	to, subject, loginLink string,
) *NotifyChangePassword {
	notify := &NotifyChangePassword{}

	notify.To = to
	notify.Subject = subject
	notify.LoginLink = loginLink

	return notify
}
