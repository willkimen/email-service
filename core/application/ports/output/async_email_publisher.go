package outputport

import "emailservice/core/application/email_request"

type AsyncEmailPublisher interface {
	Publish(request emailrequest.EmailRequest) error
}
