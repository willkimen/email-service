package outputport

// SendEmailPort defines the output port responsible for delivering
// an email to an external email provider.
//
// Implementations of this interface encapsulate the integration with
// SMTP servers or third-party email services and are responsible only
// for sending the email, not for rendering content or validating data.
type SendEmailPort interface {
	// SendEmail sends an email using the given destination, subject,
	// and rendered body content.
	//
	// Errors:
	//   -  An apperrors.InfrastructureError. Returns wrapped with apperrors.ErrTemporaryFailure
	//     on rate limit breaches, or wrapped with apperrors.ErrPermanentFailure
	//     on API or network failures.
	SendEmail(to, subject, body string) error
}
