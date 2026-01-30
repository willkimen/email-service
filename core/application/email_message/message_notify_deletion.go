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

func NewNotifyDeletion(
	to, subject string,
) *NotifyDeletion {
	notify := &NotifyDeletion{}

	notify.To = to
	notify.Subject = subject

	return notify
}
