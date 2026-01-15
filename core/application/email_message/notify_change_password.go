package emailmessage

const TemplateNotifyChangePasswordID = "notify_change_password"

type NotifyChangePassword struct {
	Base
	LoginLink  string
	TemplateId string
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
