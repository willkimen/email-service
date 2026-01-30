package emailmessage

// NotifyActivation represents an email notification sent after
// an account has been successfully activated.
//
// This message is used to inform the user that the activation
// process is complete and provides a link to access the system.
type NotifyActivation struct {
	Base

	// LoginLink defines the URL the user should access after activation.
	LoginLink string
}

// TemplateID returns the identifier of the email template
// associated with the activation notification.
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

