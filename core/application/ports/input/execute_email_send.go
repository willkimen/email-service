package inputport

import "emailservice/core/application/email_message"

type ExecuteEmailSendUseCase interface {
	ExecuteSend(message emailmessage.EmailMessage) error
}
