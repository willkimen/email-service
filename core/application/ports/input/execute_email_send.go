package inputport

import "emailservice/core/application/email_request"

type ExecuteEmailSendUseCase interface {
	ExecuteSend(request emailrequest.EmailRequest) error
}
