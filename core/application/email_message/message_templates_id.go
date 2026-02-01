package emailmessage

// Template identifiers used to select the correct email template
// when rendering and sending transactional emails.
//
// Each constant represents a specific email scenario, such as
// sending verification codes or notifying the user about completed
// account-related actions.
const (
	TemplateActivationCodeID   = "activation_code"
	TemplateNotifyActivationID = "notify_activation"

	TemplateChangeEmailCodeID   = "change_email_code"
	TemplateNotifyChangeEmailID = "notify_change_email"

	TemplateChangePasswordCodeID   = "change_password_code"
	TemplateNotifyChangePasswordID = "notify_change_password"

	TemplateResetPasswordCodeID   = "reset_password_code"
	TemplateNotifyResetPasswordID = "notify_reset_password"

	TemplateDeletionCodeID   = "deletion_code"
	TemplateNotifyDeletionID = "notify_deletion"
)
