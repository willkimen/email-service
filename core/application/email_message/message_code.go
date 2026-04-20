// This file defines concrete email message types that represent
// different verification and security-related email scenarios.
//
// Each struct composes a BaseMessage with a specific body type,
// combining shared email metadata (recipient, subject, etc.)
// with the data required to render a particular template.
//
// These message types implement the EmailMessage contract by
// exposing their EmailType identifier and the structured body
// data used during template rendering or serialization.
// EmailVerificationCode represents the email message used to deliver an email
// email verification code and related email verification details to the user.

package emailmessage

type EmailVerificationCode struct {
	BaseMessage
	EmailVerificationCodeBody
}

func (EmailVerificationCode) GetEmailType() string {
	return EmailTypeEmailVerificationCode
}

func (a *EmailVerificationCode) GetBodyData() any {
	return a.EmailVerificationCodeBody
}

// ChangeEmailCode represents the data required to send an email
// containing a verification code for changing the user's email address.
type ChangeEmailCode struct {
	BaseMessage
	ChangeEmailCodeBody
}

func (ChangeEmailCode) GetEmailType() string {
	return EmailTypeChangeEmailCode
}

func (c *ChangeEmailCode) GetBodyData() any {
	return c.ChangeEmailCodeBody
}

// DeletionCode represents the data required to send an email
// containing a verification code for account deletion operations.
type DeletionCode struct {
	BaseMessage
	DeletionCodeBody
}

func (DeletionCode) GetEmailType() string {
	return EmailTypeDeletionCode
}

func (d *DeletionCode) GetBodyData() any {
	return d.DeletionCodeBody
}

// ResetPasswordCode represents the data required to send an email
// containing a verification code and link for password reset.
type ResetPasswordCode struct {
	BaseMessage
	ResetPasswordCodeBody
}

func (ResetPasswordCode) GetEmailType() string {
	return EmailTypeResetPasswordCode
}

func (r *ResetPasswordCode) GetBodyData() any {
	return r.ResetPasswordCodeBody
}

// ChangePasswordCode represents the data required to send an email
// containing a verification code for password change operations.
type ChangePasswordCode struct {
	BaseMessage
	ChangePasswordCodeBody
}

func (ChangePasswordCode) GetEmailType() string {
	return EmailTypeChangePasswordCode
}

func (c *ChangePasswordCode) GetBodyData() any {
	return c.ChangePasswordCodeBody
}
