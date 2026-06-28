package htmlrenderer

import "emailservice/core/application/message"

// subjects maps system email event types to their respective email subject lines.
var subjects = map[string]string{
	// Account Registration & Verification
	message.MessageTypeEmailVerificationCode: "Confirm your email",
	message.MessageTypeNotifyEmailVerified:   "Email successfully verified",

	// Email Modification
	message.MessageTypeChangeEmailCode:    "Code to change your email",
	message.MessageTypeNotifyEmailChanged: "Your email has been changed",

	// Password Modification (Authenticated User)
	message.MessageTypeChangePasswordCode:    "Code to change your password",
	message.MessageTypeNotifyPasswordChanged: "Your password has been changed",

	// Account Recovery / Password Reset
	message.MessageTypeResetPasswordCode:   "Code to reset your password",
	message.MessageTypeNotifyPasswordReset: "Your password has been reset",

	// Account Deletion
	message.MessageTypeAccountDeletionCode:  "Code to delete your account",
	message.MessageTypeNotifyAccountDeleted: "Your account has been deleted",
}
