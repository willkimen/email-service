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

func (NotifyChangePassword) GetEmailType() string {
	return EmailTypeNotifyChangePassword
}

func (n *NotifyChangePassword) GetBodyData() any {
	return struct {
		LoginLink string
	}{
		LoginLink: n.LoginLink,
	}
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
