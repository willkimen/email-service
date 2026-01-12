package outputport

import "emailservice/core/application/email_request"

type EmailContentRenderer interface {
	Render(request emailrequest.EmailRequest) (body string, err error)
}
