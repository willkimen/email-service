package payload_test

const (
	to                     = "doe@email.com"
	subject                = "some subject"
	verificationCode       = "1234"
	codeExpiratinoHours    = 2
	link                   = "www.some.com/some"
	activationDeadlineDays = 7
)

// Templates
const (
	activationCodeTemplatePath   = "../templates/activation_code.html"
	notifyActivationTemplatePath = "../templates/notify_activation.html"

	changeEmailCodeTemplatePath   = "../templates/change_email_code.html"
	notifyChangeEmailTemplatePath = "../templates/notify_change_email.html"

	passwordResetCodeTemplatePath   = "../templates/password_reset_code.html"
	notifyPasswordResetTemplatePath = "../templates/notify_password_reset.html"

	deletionCodeTemplatePath   = "../templates/deletion_code.html"
	notifyDeletionTemplatePath = "../templates/notify_deletion.html"
)
