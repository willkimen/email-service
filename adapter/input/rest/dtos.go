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
// It defines the verification code and its expiration time.
type BaseCodeDTO struct {
	VerificationCode    string `json:"verification_code"`
	CodeExpirationHours string `json:"code_expiration_hours"`
}

// ========= Activation code =========

// ActivationCodeDTO represents the payload required to send
// an account activation verification code.
type ActivationCodeDTO struct {
	BaseDTO
	BaseCodeDTO
	ActivationLink         string `json:"activation_link"`
	ActivationDeadlineDays string `json:"activation_deadline_days"`
}

func (a *ActivationCodeDTO) ToEmailMessage() emailmessage.EmailMessage {
	return emailmessage.NewActivationCode(
		a.To,
		a.Subject,
		a.VerificationCode,
		a.ActivationLink,
		a.CodeExpirationHours,
		a.ActivationDeadlineDays,
	)
}

// ========= Notify activation =========

// NotifyActivationDTO represents the payload for notifying
// that an account has been successfully activated.
type NotifyActivationDTO struct {
	BaseDTO
	LoginLink string `json:"login_link"`
}

func (n *NotifyActivationDTO) ToEmailMessage() emailmessage.EmailMessage {
	return emailmessage.NewNotifiyActivation(
		n.To,
		n.Subject,
		n.LoginLink,
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
		c.CodeExpirationHours,
	)
}

// ========= Notify change email =========

// NotifyChangeEmailDTO represents the payload for notifying
// that the user's email has been changed.
type NotifyChangeEmailDTO struct {
	BaseDTO
	LoginLink string `json:"login_link"`
}

func (n *NotifyChangeEmailDTO) ToEmailMessage() emailmessage.EmailMessage {
	return emailmessage.NewNotifyChangeEmail(
		n.To,
		n.Subject,
		n.LoginLink,
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
		r.CodeExpirationHours,
	)
}

// ========= Notify change password =========

// NotifyChangePasswordDTO represents the payload for notifying
// that the user's password has been changed.
type NotifyChangePasswordDTO struct {
	BaseDTO
	LoginLink string `json:"login_link"`
}

func (n *NotifyChangePasswordDTO) ToEmailMessage() emailmessage.EmailMessage {
	return emailmessage.NewNotifyChangePassword(
		n.To,
		n.Subject,
		n.LoginLink,
	)
}

// ========= Reset password code =========

// ResetPasswordCodeDTO represents the payload for sending
// a password reset verification code.
type ResetPasswordCodeDTO struct {
	BaseDTO
	BaseCodeDTO
	ResetPasswordLink string `json:"reset_password_link"`
}

func (r *ResetPasswordCodeDTO) ToEmailMessage() emailmessage.EmailMessage {
	return emailmessage.NewResetPasswordCode(
		r.To,
		r.Subject,
		r.VerificationCode,
		r.CodeExpirationHours,
		r.ResetPasswordLink,
	)
}

// ========= Notify reset password =========

// NotifyResetPasswordDTO represents the payload for notifying
// that the user's password has been reset.
type NotifyResetPasswordDTO struct {
	BaseDTO
	LoginLink string `json:"login_link"`
}

func (n *NotifyResetPasswordDTO) ToEmailMessage() emailmessage.EmailMessage {
	return emailmessage.NewNotifyResetPassword(
		n.To,
		n.Subject,
		n.LoginLink,
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
		d.CodeExpirationHours,
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
