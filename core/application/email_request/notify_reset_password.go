package emailrequest

const TemplateNotifyResetPasswordID = "notify_reset_password"

type NotifyResetPassword struct {
	Base
	LoginLink string
}

func (NotifyResetPassword) TemplateID() string {
	return TemplateNotifyResetPasswordID
}

func NewNotifyResetPassword(
	to, subject, loginLink string,
) *NotifyResetPassword {
	notify := &NotifyResetPassword{}

	notify.To = to
	notify.Subject = subject
	notify.LoginLink = loginLink

	return notify
}
