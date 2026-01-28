package emailmessage

type NotifyDeletion struct {
	Base
}

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
