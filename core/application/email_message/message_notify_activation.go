package emailmessage

// NotifyActivation represents an email notification sent after
// an account has been successfully activated.
//
// This message is used to inform the user that the activation
// process is complete and provides a link to access the system.
type NotifyActivation struct {
	BaseMessage

	// LoginLink defines the URL the user should access after activation.
	LoginLink string
}

func (NotifyActivation) GetEmailType() string {
	return EmailTypeNotifyActivation
}

func (n *NotifyActivation) GetBodyData() any {
	return struct {
		LoginLink string
	}{
		LoginLink: n.LoginLink,
	}
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
