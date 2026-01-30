package emailmessage

// Template identifiers used to select the correct email template
// when rendering and sending transactional emails.
//
// Each constant represents a specific email scenario, such as
// sending verification codes or notifying the user about completed
// account-related actions.
const (
	TemplateActivationCodeID     = "activation_code"
	TemplateChangeEmailCodeID    = "change_email_code"
	TemplateChangePasswordCodeID = "change_password_code"
	TemplateResetPasswordCodeID  = "reset_password_code"
	TemplateDeletionCodeID       = "deletion_code"

	TemplateNotifyActivationID     = "notify_activation"
	TemplateNotifyChangeEmailID    = "notify_change_email"
	TemplateNotifyChangePasswordID = "notify_change_password"
	TemplateNotifyResetPasswordID  = "notify_reset_password"
	TemplateNotifyDeletionID       = "notify_deletion"
)
