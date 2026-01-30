package emailmessage

import (
	"regexp"
	"strings"
)

// Validation rules for all email message types are centralized in this file.
//
// The validation model is based on composition:
// - Base validates fields common to all email messages.
// - BaseCode validates fields common to code-based emails.
// - Concrete message types compose these structures and add only
//   the rules specific to their own data.
//
// Each ValidateData method enforces allowed states and returns
// validation errors when constraints are violated.

const (
	FieldTo                     = "to"
	FieldSubject                = "subject"
	FieldVerificationCode       = "verificationCode"
	FieldCodeExpirationHours    = "codeExpirationHours"
	FieldActivationLink         = "activationLink"
	FieldActivationDeadlineDays = "activationDeadlineDays"
	FieldResetPasswordLink      = "resetPasswordLink"
	FieldLoginLink              = "loginLink"
)

// fieldRule binds a raw value to its logical field name,
// allowing validation errors to report which field is invalid.
type fieldRule struct {
	value string
	name  string
}

// validatorEmailFormat ensures the email address follows a valid format.
// Only syntactically valid email addresses are accepted.
func validatorEmailFormat(email string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return NewEmailInvalidFormatError()
	}
	return nil
}

// validateRequiredFields ensures that all provided fields contain data.
// Empty or whitespace-only values are not accepted.
func validateRequiredFields(fields ...fieldRule) error {
	for _, f := range fields {
		if strings.TrimSpace(f.value) == "" {
			return NewEmptyFieldError(f.name)
		}
	}
	return nil
}

// ValidateData validates the minimal required state for any email message.
// All messages must define a recipient and a subject, and the email
// address must be syntactically valid.
func (p *Base) ValidateData() error {
	if err := validateRequiredFields(
		fieldRule{p.To, FieldTo},
		fieldRule{p.Subject, FieldSubject},
	); err != nil {
		return err
	}

	if err := validatorEmailFormat(p.To); err != nil {
		return err
	}

	return nil
}

// ValidateData validates fields common to code-based emails.
// Code-based messages must define a verification code and its expiration.
func (p *BaseCode) ValidateData() error {
	if err := validateRequiredFields(
		fieldRule{p.VerificationCode, FieldVerificationCode},
		fieldRule{p.CodeExpirationHours, FieldCodeExpirationHours},
	); err != nil {
		return err
	}

	return nil
}

// ValidateData validates activation code emails.
// In addition to base validations, activation-specific data must be present.
func (p *ActivationCode) ValidateData() error {
	if err := p.Base.ValidateData(); err != nil {
		return err
	}

	if err := p.BaseCode.ValidateData(); err != nil {
		return err
	}

	if err := validateRequiredFields(
		fieldRule{p.ActivationLink, FieldActivationLink},
		fieldRule{p.ActivationDeadlineDays, FieldActivationDeadlineDays},
	); err != nil {
		return err
	}

	return nil
}

// ValidateData validates change email code messages.
// No additional fields are required beyond Base and BaseCode.
func (p *ChangeEmailCode) ValidateData() error {
	if err := p.Base.ValidateData(); err != nil {
		return err
	}

	if err := p.BaseCode.ValidateData(); err != nil {
		return err
	}

	return nil
}

// ValidateData validates reset password code messages.
// A reset password link must be provided.
func (p *ResetPasswordCode) ValidateData() error {
	if err := p.Base.ValidateData(); err != nil {
		return err
	}

	if err := p.BaseCode.ValidateData(); err != nil {
		return err
	}

	if err := validateRequiredFields(
		fieldRule{p.ResetPasswordLink, FieldResetPasswordLink},
	); err != nil {
		return err
	}

	return nil
}

// ValidateData validates account deletion code messages.
// Only base email and code-related fields are required.
func (p *DeletionCode) ValidateData() error {
	if err := p.Base.ValidateData(); err != nil {
		return err
	}

	if err := p.BaseCode.ValidateData(); err != nil {
		return err
	}

	return nil
}

// ValidateData validates change password code messages.
// No additional fields beyond Base and BaseCode are required.
func (p *ChangePasswordCode) ValidateData() error {
	if err := p.Base.ValidateData(); err != nil {
		return err
	}

	if err := p.BaseCode.ValidateData(); err != nil {
		return err
	}

	return nil
}

// ValidateData validates activation notification emails.
// A login link must be present to allow user access.
func (p *NotifyActivation) ValidateData() error {
	if err := p.Base.ValidateData(); err != nil {
		return err
	}

	if err := validateRequiredFields(
		fieldRule{p.LoginLink, FieldLoginLink},
	); err != nil {
		return err
	}

	return nil
}

// ValidateData validates change email notification emails.
// A login link is required to allow user confirmation.
func (p *NotifyChangeEmail) ValidateData() error {
	if err := p.Base.ValidateData(); err != nil {
		return err
	}

	if err := validateRequiredFields(
		fieldRule{p.LoginLink, FieldLoginLink},
	); err != nil {
		return err
	}

	return nil
}

// ValidateData validates reset password notification emails.
// A login link must be provided.
func (p *NotifyResetPassword) ValidateData() error {
	if err := p.Base.ValidateData(); err != nil {
		return err
	}

	if err := validateRequiredFields(
		fieldRule{p.LoginLink, FieldLoginLink},
	); err != nil {
		return err
	}

	return nil
}

// ValidateData validates account deletion notification emails.
// Only base email fields are required.
func (p *NotifyDeletion) ValidateData() error {
	if err := p.Base.ValidateData(); err != nil {
		return err
	}

	return nil
}

// ValidateData validates change password notification emails.
// A login link is required to allow user access.
func (p *NotifyChangePassword) ValidateData() error {
	if err := p.Base.ValidateData(); err != nil {
		return err
	}

	if err := validateRequiredFields(
		fieldRule{p.LoginLink, FieldLoginLink},
	); err != nil {
		return err
	}

	return nil
}

