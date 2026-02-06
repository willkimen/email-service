package emailmessage

// NotifyDeletion represents an email notification sent after
// a user's account has been successfully deleted.
//
// This message is used to inform the user that the deletion
// process has been completed.
type NotifyDeletion struct {
	BaseMessage
}

// TemplateID returns the identifier of the email template
// associated with the account deletion notification.
func (NotifyDeletion) TemplateID() string {
	return TemplateNotifyDeletionID
}

// GetBodyData returns the data structure used to populate the email template
// associated with the entity.
//
// The returned value contains only the fields required by the template
// renderer and represents a read-only projection of the entity.
// This method does not apply formatting or validation logic;
// it simply exposes the data needed for template interpolation.
func (NotifyDeletion) GetBodyData() any {
	// In this specific case, there is no data to be rendered in the email body.
	return nil
}

func NewNotifyDeletion(
	to, subject string,
) *NotifyDeletion {
	notify := &NotifyDeletion{}

	notify.To = to
	notify.Subject = subject

	return notify
}
