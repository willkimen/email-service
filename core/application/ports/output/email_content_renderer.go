package outputport

import "emailservice/core/application/email_message"

type EmailContentRenderer interface {
	Render(message emailmessage.EmailMessage) (string, error)
}
