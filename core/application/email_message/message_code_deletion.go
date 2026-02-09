package emailmessage

// DeletionCode represents the data required to send an email
// containing a verification code for account deletion operations.
type DeletionCode struct {
	BaseMessage
	BaseCodeMessage
}

func (DeletionCode) GetEmailType() string {
	return EmailTypeDeletionCode
}

func (d *DeletionCode) GetBodyData() any {
	return struct {
		VerificationCode    string
		CodeExpirationHours string
	}{
		VerificationCode:    d.VerificationCode,
		CodeExpirationHours: d.CodeExpirationHours,
	}
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
