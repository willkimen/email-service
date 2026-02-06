package outputport

// SendEmailOutputPort defines the output port responsible for delivering
// an email to an external email provider.
//
// Implementations of this interface encapsulate the integration with
// SMTP servers or third-party email services and are responsible only
// for sending the email, not for rendering content or validating data.
type SendEmailOutputPort interface {
	// SendEmail sends an email using the given destination, subject,
	// and rendered body content.
	//
	// It returns an error if the email cannot be delivered.
	SendEmail(to, subject, body string) error
}
