package emailmessage

// EmailType identifiers represent the different types of transactional
// emails supported by the system.
//
// These constants describe the semantic intent of each email, not the
// rendering mechanism itself. They are used across the codebase to:
//
//   - Identify which type of email is being sent
//   - Drive business decisions in adapters and services
//   - Select the appropriate HTML template during rendering
//
// Using explicit email types avoids hardcoded strings, improves
// type safety, and provides a shared vocabulary between layers.
const (
	EmailTypeActivationCode   = "activation_code"
	EmailTypeNotifyActivation = "notify_activation"

	EmailTypeChangeEmailCode   = "change_email_code"
	EmailTypeNotifyChangeEmail = "notify_change_email"

	EmailTypeChangePasswordCode   = "change_password_code"
	EmailTypeNotifyChangePassword = "notify_change_password"

	EmailTypeResetPasswordCode   = "reset_password_code"
	EmailTypeNotifyResetPassword = "notify_reset_password"

	EmailTypeDeletionCode   = "deletion_code"
	EmailTypeNotifyDeletion = "notify_deletion"
)
