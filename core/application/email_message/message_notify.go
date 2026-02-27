// This file defines notification-type email messages that are sent
// after a security-sensitive action has been successfully completed.
//
// Unlike code-based messages, these notifications do not request
// further verification. They exist to inform the user that an
// operation (such as activation, password change, or deletion)
// has already been finalized.
//
// Each message composes BaseMessage with a specific body structure
// and implements the EmailMessage contract by providing its type
// identifier and template body data.
// NotifyActivation represents an email notification sent after
// an account has been successfully activated.

package emailmessage

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

// NotifyChangeEmail represents an email notification sent after
// a user's email address has been successfully changed.
//
// This message informs the user about the change and provides
// a link to access the system.
type NotifyChangeEmail struct {
	BaseMessage
	NotifyChangeEmailBody
}

func (NotifyChangeEmail) GetEmailType() string {
	return EmailTypeNotifyChangeEmail
}

func (n *NotifyChangeEmail) GetBodyData() any {
	return n.NotifyChangeEmailBody
}

// NotifyChangePassword represents an email notification sent after
// a user's password has been successfully changed.
//
// This message is used to inform the user about the password change
// and provide a link to access the system.
type NotifyChangePassword struct {
	BaseMessage
	NotifyChangePasswordBody
}

func (NotifyChangePassword) GetEmailType() string {
	return EmailTypeNotifyChangePassword
}

func (n *NotifyChangePassword) GetBodyData() any {
	return n.NotifyChangePasswordBody
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

// NotifyResetPassword represents an email notification sent after
// a password reset has been completed.
//
// This message is used to inform the user that the password
// was successfully changed and provides a link for login.
type NotifyResetPassword struct {
	BaseMessage
	NotifyResetPasswordBody
}

func (NotifyResetPassword) GetEmailType() string {
	return EmailTypeNotifyResetPassword
}

func (n *NotifyResetPassword) GetBodyData() any {
	return n.NotifyResetPasswordBody
}
