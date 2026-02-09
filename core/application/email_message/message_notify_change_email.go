package emailmessage

// NotifyChangeEmail represents an email notification sent after
// a user's email address has been successfully changed.
//
// This message informs the user about the change and provides
// a link to access the system.
type NotifyChangeEmail struct {
	BaseMessage

	// LoginLink defines the URL the user should use to access the system
	// after the email change is completed.
	LoginLink string
}

func (NotifyChangeEmail) GetEmailType() string {
	return EmailTypeNotifyChangeEmail
}

func (n *NotifyChangeEmail) GetBodyData() any {
	return struct {
		LoginLink string
	}{
		LoginLink: n.LoginLink,
	}
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
