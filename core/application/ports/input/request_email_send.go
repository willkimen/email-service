package inputport

import "emailservice/core/application/email_message"

type RequestEmailSendUseCase interface {
	Request(message emailmessage.EmailMessage) error
}
