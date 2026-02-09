package emailmessage

type NotifyDeletionBody struct {
}

// NotifyDeletion represents an email notification sent after
// a user's account has been successfully deleted.
//
// This message is used to inform the user that the deletion
// process has been completed.
type NotifyDeletion struct {
	BaseMessage
	NotifyDeletionBody
}

func (NotifyDeletion) GetEmailType() string {
	return EmailTypeNotifyDeletion
}

func (n *NotifyDeletion) GetBodyData() any {
	return n.NotifyDeletionBody
}

func NewNotifyDeletion(
	to, subject string,
) *NotifyDeletion {
	notify := &NotifyDeletion{}

	notify.To = to
	notify.Subject = subject

	return notify
}
