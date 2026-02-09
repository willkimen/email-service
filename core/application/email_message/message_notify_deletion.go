package emailmessage

// NotifyDeletion represents an email notification sent after
// a user's account has been successfully deleted.
//
// This message is used to inform the user that the deletion
// process has been completed.
type NotifyDeletion struct {
	BaseMessage
}

func (NotifyDeletion) GetEmailType() string {
	return EmailTypeNotifyDeletion
}

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
