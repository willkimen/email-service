package inputport

import "emailservice/core/application/email_request"

type RequestEmailSendUseCase interface {
	RequestSend(request emailrequest.EmailRequest) error
}
