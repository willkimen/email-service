package outputport

import "emailservice/core/application/email_message"

type EmailRequestPublisher interface {
	Publish(message emailmessage.EmailMessage) error
}
