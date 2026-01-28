package rest

import emailmessage "emailservice/core/application/email_message"

type EmailRequestDTO interface {
	ToEmailMessage() emailmessage.EmailMessage
}

// ========= Bases =========
type BaseDTO struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
}

type BaseCodeDTO struct {
	VerificationCode    string `json:"verification_code"`
	CodeExpirationHours string `json:"code_expiration_hours"`
}

// ========= Activation code =========
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

// =========Notify activation =========
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

// ========= Notify reset password code =========
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
type NotifyDeletionDTO struct {
	BaseDTO
}

func (n *NotifyDeletionDTO) ToEmailMessage() emailmessage.EmailMessage {
	return emailmessage.NewNotifyDeletion(
		n.To,
		n.Subject,
	)
}
