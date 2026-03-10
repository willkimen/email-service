package renderer

import "emailservice/core/application/email_message"

// pathTemplates maps email types to their corresponding
// HTML template file paths.
//
// This map defines the relationship between an EmailType
// (a concept that represents a kind of email)
// and the physical HTML template used to render it.
//
// The renderer uses this map to resolve which HTML template
// should be loaded and executed for a given email type.
//
// To add a new email template:
// 1. Create the HTML file under the templates directory.
// 2. Define a new EmailType constant.
// 3. Register the EmailType and its template path in this map.
var pathTemplates = map[string]string{
	emailmessage.EmailTypeActivationCode:   "templates/activation_code.html",
	emailmessage.EmailTypeNotifyActivation: "templates/notify_activation.html",

	emailmessage.EmailTypeChangeEmailCode:   "templates/change_email_code.html",
	emailmessage.EmailTypeNotifyChangeEmail: "templates/notify_change_email.html",

	emailmessage.EmailTypeChangePasswordCode:   "templates/change_password_code.html",
	emailmessage.EmailTypeNotifyChangePassword: "templates/notify_change_password.html",

	emailmessage.EmailTypeResetPasswordCode:   "templates/reset_password_code.html",
	emailmessage.EmailTypeNotifyResetPassword: "templates/notify_reset_password.html",

	emailmessage.EmailTypeDeletionCode:   "templates/deletion_code.html",
	emailmessage.EmailTypeNotifyDeletion: "templates/notify_deletion.html",
}
