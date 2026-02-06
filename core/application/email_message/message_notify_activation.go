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

// TemplateID returns the identifier of the email template
// associated with the activation notification.
func (NotifyActivation) TemplateID() string {
	return TemplateNotifyActivationID
}

// GetBodyData returns the data structure used to populate the email template
// associated with the entity.
//
// The returned value contains only the fields required by the template
// renderer and represents a read-only projection of the entity.
// This method does not apply formatting or validation logic;
// it simply exposes the data needed for template interpolation.
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
