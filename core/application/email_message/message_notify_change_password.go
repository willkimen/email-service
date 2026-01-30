package emailmessage

// NotifyChangePassword represents an email notification sent after
// a user's password has been successfully changed.
//
// This message is used to inform the user about the password change
// and provide a link to access the system.
type NotifyChangePassword struct {
	BaseMessage

	// LoginLink defines the URL the user should use to access the system
	// after the password change is completed.
	LoginLink string
}

// TemplateID returns the identifier of the email template
// associated with the password change notification.
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
