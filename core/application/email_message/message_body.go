// Package emailmessage defines the body structures used by different
// email message types within the application.
//
// Each body struct represents the data required to populate a specific
// email template. These structures are intentionally lightweight and
// contain only the fields necessary for template rendering.
//
// The bodies are part of the application layer and serve as structured
// data carriers between the domain message types and infrastructure
// components such as template renderers or background task serializers.
//
// Some body types embed shared structures (e.g., BaseCodeMessage)
// to reuse common fields like verification codes and expiration metadata,
// ensuring consistency across email variants.
//
// These types do not contain behavior. They exist solely to model
// the template input data associated with each email scenario.

package emailmessage

type EmailVerificationCodeBody struct {
	BaseCodeMessage

	// EmailVerificationDeadlineDays defines how many days the email verification remains valid.
	EmailVerificationDeadlineDays string
}

type ChangeEmailCodeBody struct {
	BaseCodeMessage
}

type ChangePasswordCodeBody struct {
	BaseCodeMessage
}

type DeletionCodeBody struct {
	BaseCodeMessage
}

type ResetPasswordCodeBody struct {
	BaseCodeMessage
}

type NotifyEmailVerificationBody struct {}

type NotifyChangeEmailBody struct {}

type NotifyChangePasswordBody struct {}

type NotifyDeletionBody struct {}

type NotifyResetPasswordBody struct {}
