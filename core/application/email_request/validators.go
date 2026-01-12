package emailrequest

import (
	"regexp"
	"strings"
)

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

type fieldRule struct {
	value string
	name  string
}

func validatorEmailFormat(email string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return NewEmailInvalidFormatError()
	}
	return nil
}

func validateRequiredFields(fields ...fieldRule) error {
	for _, f := range fields {
		if strings.TrimSpace(f.value) == "" {
			return NewEmptyFieldError(f.name)
		}
	}
	return nil
}

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

func (p *BaseCode) ValidateData() error {
	if err := validateRequiredFields(
		fieldRule{p.VerificationCode, FieldVerificationCode},
		fieldRule{p.CodeExpirationHours, FieldCodeExpirationHours},
	); err != nil {
		return err
	}

	return nil
}

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

func (p *ChangeEmailCode) ValidateData() error {
	if err := p.Base.ValidateData(); err != nil {
		return err
	}

	if err := p.BaseCode.ValidateData(); err != nil {
		return err
	}

	return nil
}

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

func (p *DeletionCode) ValidateData() error {
	if err := p.Base.ValidateData(); err != nil {
		return err
	}

	if err := p.BaseCode.ValidateData(); err != nil {
		return err
	}

	return nil
}

func (p *ChangePasswordCode) ValidateData() error {
	if err := p.Base.ValidateData(); err != nil {
		return err
	}

	if err := p.BaseCode.ValidateData(); err != nil {
		return err
	}

	return nil
}

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

func (p *NotifyDeletion) ValidateData() error {
	if err := p.Base.ValidateData(); err != nil {
		return err
	}

	return nil
}

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
