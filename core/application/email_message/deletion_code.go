package emailmessage

const TemplateDeletionCodeID = "deletion_code"

type DeletionCode struct {
	Base
	BaseCode
}

func (DeletionCode) TemplateID() string {
	return TemplateDeletionCodeID
}

func NewDeletionCode(
	to, subject, verificationCode, codeExpirationHours string,
) *DeletionCode {
	deletion := &DeletionCode{}

	deletion.To = to
	deletion.Subject = subject
	deletion.VerificationCode = verificationCode
	deletion.CodeExpirationHours = codeExpirationHours

	return deletion
}
