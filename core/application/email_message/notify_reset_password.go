package emailmessage

// NotifyResetPassword represents an email notification sent after
// a password reset has been completed.
//
// This message is used to inform the user that the password
// was successfully changed and provides a link for login.
type NotifyResetPassword struct {
	Base

	// LoginLink defines the URL the user can use to access
	// the application after resetting the password.
	LoginLink string
}

// TemplateID returns the identifier of the email template
// associated with the password reset notification.
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

