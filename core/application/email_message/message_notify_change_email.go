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

// TemplateID returns the identifier of the email template
// associated with the change email notification.
func (NotifyChangeEmail) TemplateID() string {
	return TemplateNotifyChangeEmailID
}

// GetBodyData returns the data structure used to populate the email template
// associated with the entity.
//
// The returned value contains only the fields required by the template
// renderer and represents a read-only projection of the entity.
// This method does not apply formatting or validation logic;
// it simply exposes the data needed for template interpolation.
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
