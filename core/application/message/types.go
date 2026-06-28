package message

// MessageType identifiers represent the different types of transactional
// emails supported by the system.
//
// These constants describe the semantic intent of each email, not the
// rendering mechanism itself. They are used across the codebase to:
//
//   - Identify which type of email is being sent
//   - Drive business decisions in adapters and services
//   - Select the appropriate HTML template during rendering
//
// Using explicit message types avoids hardcoded strings, improves
// type safety, and provides a shared vocabulary between layers.
const (
	MessageTypeEmailVerificationCode = "email_verification_code"
	MessageTypeNotifyEmailVerified   = "notify_email_verified"

	MessageTypeChangeEmailCode    = "change_email_code"
	MessageTypeNotifyEmailChanged = "notify_email_changed"

	MessageTypeChangePasswordCode    = "change_password_code"
	MessageTypeNotifyPasswordChanged = "notify_password_changed"

	MessageTypeResetPasswordCode   = "reset_password_code"
	MessageTypeNotifyPasswordReset = "notify_password_reset"

	MessageTypeAccountDeletionCode  = "account_deletion_code"
	MessageTypeNotifyAccountDeleted = "notify_account_deleted"
)
