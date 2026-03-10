package emailmessage

import (
	"regexp"
	"strings"
)

// Validation rules for all email message types are centralized in this file.
//
// The validation model is based on composition:
// - BaseMessage validates fields common to all email messages.
// - BaseCodeMessage validates fields common to code-based emails.
// - Concrete message types compose these structures and add only
//   the rules specific to their own data.
//
// Each ValidateData method enforces allowed states and returns
// validation errors when constraints are violated.

const (
	FieldTo                     = "to"
	FieldSubject                = "subject"
	FieldVerificationCode       = "verification_code"
	FieldCodeExpirationHours    = "code_expiration_hours"
	FieldActivationLink         = "activation_link"
	FieldActivationDeadlineDays = "activation_deadline_days"
	FieldResetPasswordLink      = "reset_password_link"
	FieldLoginLink              = "login_link"
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

func (b *BaseMessage) ValidateData() error {
	if err := validateRequiredFields(
		fieldRule{b.To, FieldTo},
		fieldRule{b.Subject, FieldSubject},
	); err != nil {
		return err
	}

	if err := validatorEmailFormat(b.To); err != nil {
		return err
	}

	return nil
}

func (b *BaseCodeMessage) ValidateData() error {
	if err := validateRequiredFields(
		fieldRule{b.VerificationCode, FieldVerificationCode},
		fieldRule{b.CodeExpirationHours, FieldCodeExpirationHours},
	); err != nil {
		return err
	}

	return nil
}

func (a *ActivationCode) ValidateData() error {
	if err := a.BaseMessage.ValidateData(); err != nil {
		return err
	}

	if err := a.BaseCodeMessage.ValidateData(); err != nil {
		return err
	}

	if err := validateRequiredFields(
		fieldRule{a.ActivationLink, FieldActivationLink},
		fieldRule{a.ActivationDeadlineDays, FieldActivationDeadlineDays},
	); err != nil {
		return err
	}

	return nil
}

func (c *ChangeEmailCode) ValidateData() error {
	if err := c.BaseMessage.ValidateData(); err != nil {
		return err
	}

	if err := c.BaseCodeMessage.ValidateData(); err != nil {
		return err
	}

	return nil
}

func (r *ResetPasswordCode) ValidateData() error {
	if err := r.BaseMessage.ValidateData(); err != nil {
		return err
	}

	if err := r.BaseCodeMessage.ValidateData(); err != nil {
		return err
	}

	if err := validateRequiredFields(
		fieldRule{r.ResetPasswordLink, FieldResetPasswordLink},
	); err != nil {
		return err
	}

	return nil
}

func (d *DeletionCode) ValidateData() error {
	if err := d.BaseMessage.ValidateData(); err != nil {
		return err
	}

	if err := d.BaseCodeMessage.ValidateData(); err != nil {
		return err
	}

	return nil
}

func (c *ChangePasswordCode) ValidateData() error {
	if err := c.BaseMessage.ValidateData(); err != nil {
		return err
	}

	if err := c.BaseCodeMessage.ValidateData(); err != nil {
		return err
	}

	return nil
}

func (p *NotifyActivation) ValidateData() error {
	if err := p.BaseMessage.ValidateData(); err != nil {
		return err
	}

	if err := validateRequiredFields(
		fieldRule{p.LoginLink, FieldLoginLink},
	); err != nil {
		return err
	}

	return nil
}

func (n *NotifyChangeEmail) ValidateData() error {
	if err := n.BaseMessage.ValidateData(); err != nil {
		return err
	}

	if err := validateRequiredFields(
		fieldRule{n.LoginLink, FieldLoginLink},
	); err != nil {
		return err
	}

	return nil
}

func (n *NotifyResetPassword) ValidateData() error {
	if err := n.BaseMessage.ValidateData(); err != nil {
		return err
	}

	if err := validateRequiredFields(
		fieldRule{n.LoginLink, FieldLoginLink},
	); err != nil {
		return err
	}

	return nil
}

func (n *NotifyDeletion) ValidateData() error {
	if err := n.BaseMessage.ValidateData(); err != nil {
		return err
	}

	return nil
}

func (n *NotifyChangePassword) ValidateData() error {
	if err := n.BaseMessage.ValidateData(); err != nil {
		return err
	}

	if err := validateRequiredFields(
		fieldRule{n.LoginLink, FieldLoginLink},
	); err != nil {
		return err
	}

	return nil
}
