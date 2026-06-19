package rest

import emailmessage "emailservice/core/application/email_message"

// EmailRequestDTO defines the contract for HTTP DTOs that can be
// converted into an email message.
//
// Each DTO represents the external HTTP payload and is responsible
// for converting itself into the corresponding representation.
type EmailRequestDTO interface {
	ToEmailMessage() emailmessage.EmailMessage
}

// ========= Bases =========

// BaseDTO represents common fields shared by all email requests.
// It defines the recipient and the email subject.
type BaseDTO struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
}

// BaseCodeDTO represents common fields used by verification code emails.
type BaseCodeDTO struct {
	VerificationCode    string `json:"verification_code"`
}

// ========= Email verification code =========

// EmailVerificationCodeDTO represents the payload required to send
// an email verification code.
type EmailVerificationCodeDTO struct {
	BaseDTO
	BaseCodeDTO
}

func (a *EmailVerificationCodeDTO) ToEmailMessage() emailmessage.EmailMessage {
	return emailmessage.NewEmailVerificationCode(
		a.To,
		a.Subject,
		a.VerificationCode,
	)
}

// ========= Notify email verification =========

// NotifyEmailVerificationDTO represents the payload for notifying
// that an email has been successfully verified.
type NotifyEmailVerificationDTO struct {
	BaseDTO
}

func (n *NotifyEmailVerificationDTO) ToEmailMessage() emailmessage.EmailMessage {
	return emailmessage.NewNotifyEmailVerification(
		n.To,
		n.Subject,
	)
}

// ========= Change email code =========

// ChangeEmailCodeDTO represents the payload for sending
// a verification code to confirm an email change.
type ChangeEmailCodeDTO struct {
	BaseDTO
	BaseCodeDTO
}

func (c *ChangeEmailCodeDTO) ToEmailMessage() emailmessage.EmailMessage {
	return emailmessage.NewChangeEmailCode(
		c.To,
		c.Subject,
		c.VerificationCode,
	)
}

// ========= Notify change email =========

// NotifyChangeEmailDTO represents the payload for notifying
// that the user's email has been changed.
type NotifyChangeEmailDTO struct {
	BaseDTO
}

func (n *NotifyChangeEmailDTO) ToEmailMessage() emailmessage.EmailMessage {
	return emailmessage.NewNotifyChangeEmail(
		n.To,
		n.Subject,
	)
}

// ========= Change password code =========

// ChangePasswordCodeDTO represents the payload for sending
// a verification code to confirm a password change.
type ChangePasswordCodeDTO struct {
	BaseDTO
	BaseCodeDTO
}

func (r *ChangePasswordCodeDTO) ToEmailMessage() emailmessage.EmailMessage {
	return emailmessage.NewChangePasswordCode(
		r.To,
		r.Subject,
		r.VerificationCode,
	)
}

// ========= Notify change password =========

// NotifyChangePasswordDTO represents the payload for notifying
// that the user's password has been changed.
type NotifyChangePasswordDTO struct {
	BaseDTO
}

func (n *NotifyChangePasswordDTO) ToEmailMessage() emailmessage.EmailMessage {
	return emailmessage.NewNotifyChangePassword(
		n.To,
		n.Subject,
	)
}

// ========= Reset password code =========

// ResetPasswordCodeDTO represents the payload for sending
// a password reset verification code.
type ResetPasswordCodeDTO struct {
	BaseDTO
	BaseCodeDTO
}

func (r *ResetPasswordCodeDTO) ToEmailMessage() emailmessage.EmailMessage {
	return emailmessage.NewResetPasswordCode(
		r.To,
		r.Subject,
		r.VerificationCode,
	)
}

// ========= Notify reset password =========

// NotifyResetPasswordDTO represents the payload for notifying
// that the user's password has been reset.
type NotifyResetPasswordDTO struct {
	BaseDTO
}

func (n *NotifyResetPasswordDTO) ToEmailMessage() emailmessage.EmailMessage {
	return emailmessage.NewNotifyResetPassword(
		n.To,
		n.Subject,
	)
}

// ========= Deletion code =========

// DeletionCodeDTO represents the payload for sending
// a verification code to confirm account deletion.
type DeletionCodeDTO struct {
	BaseDTO
	BaseCodeDTO
}

func (d *DeletionCodeDTO) ToEmailMessage() emailmessage.EmailMessage {
	return emailmessage.NewDeletionCode(
		d.To,
		d.Subject,
		d.VerificationCode,
	)
}

// ========= Notify deletion =========

// NotifyDeletionDTO represents the payload for notifying
// that the user's account has been deleted.
type NotifyDeletionDTO struct {
	BaseDTO
}

func (n *NotifyDeletionDTO) ToEmailMessage() emailmessage.EmailMessage {
	return emailmessage.NewNotifyDeletion(
		n.To,
		n.Subject,
	)
}
