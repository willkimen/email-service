package emailmessage

type NotifyActivationBody struct {
	// LoginLink defines the URL the user should access after activation.
	LoginLink string
}

// NotifyActivation represents an email notification sent after
// an account has been successfully activated.
//
// This message is used to inform the user that the activation
// process is complete and provides a link to access the system.
type NotifyActivation struct {
	BaseMessage
	NotifyActivationBody
}

func (NotifyActivation) GetEmailType() string {
	return EmailTypeNotifyActivation
}

func (n *NotifyActivation) GetBodyData() any {
	return n.NotifyActivationBody
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
