package htmlrenderer

import "emailservice/core/application/message"

// pathTemplates maps message types to their corresponding
// HTML template file paths.
//
// This map defines the relationship between an MessageType
// (a concept that represents a kind of email)
// and the physical HTML template used to render it.
//
// The renderer uses this map to resolve which HTML template
// should be loaded and executed for a given message type.
//
// To add a new email template:
// 1. Create the HTML file under the templates directory.
// 2. Define a new MessageType constant.
// 3. Register the MessageType and its template path in this map.
var pathTemplates = map[string]string{
	message.MessageTypeEmailVerificationCode: "templates/email_verification_code.html",
	message.MessageTypeNotifyEmailVerified:   "templates/notify_email_verified.html",

	message.MessageTypeChangeEmailCode:    "templates/change_email_code.html",
	message.MessageTypeNotifyEmailChanged: "templates/notify_email_changed.html",

	message.MessageTypeChangePasswordCode:    "templates/change_password_code.html",
	message.MessageTypeNotifyPasswordChanged: "templates/notify_password_changed.html",

	message.MessageTypeResetPasswordCode:   "templates/reset_password_code.html",
	message.MessageTypeNotifyPasswordReset: "templates/notify_password_reset.html",

	message.MessageTypeAccountDeletionCode:  "templates/account_deletion_code.html",
	message.MessageTypeNotifyAccountDeleted: "templates/notify_account_deleted.html",
}
