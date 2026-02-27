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

type ActivationCodeBody struct {
	BaseCodeMessage
	// ActivationLink is the URL the user must access to activate the account.
	ActivationLink string

	// ActivationDeadlineDays defines how many days the activation remains valid.
	ActivationDeadlineDays string
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

	// ResetPasswordLink is the URL the user must access to complete
	// the password reset process.
	ResetPasswordLink string
}

type NotifyActivationBody struct {
	// LoginLink defines the URL the user should access after activation.
	LoginLink string
}

type NotifyChangeEmailBody struct {
	// LoginLink defines the URL the user should use to access the system
	// after the email change is completed.
	LoginLink string
}

type NotifyChangePasswordBody struct {
	// LoginLink defines the URL the user should use to access the system
	// after the password change is completed.
	LoginLink string
}

type NotifyDeletionBody struct {
}

type NotifyResetPasswordBody struct {
	// LoginLink defines the URL the user can use to access
	// the application after resetting the password.
	LoginLink string
}
