package emailresend_test

import (
	"errors"

	"github.com/resend/resend-go/v3"
)

// mockResendEmailAPI is a test double used to simulate
// different behaviors of the Resend client.
type mockResendEmailAPI struct {
	sendFunc func(params *resend.SendEmailRequest) (*resend.SendEmailResponse, error)
}

func (m *mockResendEmailAPI) Send(
	params *resend.SendEmailRequest,
) (*resend.SendEmailResponse, error) {
	return m.sendFunc(params)
}

var mockEmailsSuccess = &mockResendEmailAPI{
	sendFunc: func(params *resend.SendEmailRequest) (*resend.SendEmailResponse, error) {
		return &resend.SendEmailResponse{}, nil
	},
}

var mockEmailsWithTemporaryFailure = &mockResendEmailAPI{
	sendFunc: func(params *resend.SendEmailRequest) (*resend.SendEmailResponse, error) {
		return nil, resend.ErrRateLimit
	},
}

var errCausePermanentFailure = errors.New("some permanent failure")
var mockEmailsWithPermanentFailure = &mockResendEmailAPI{
	sendFunc: func(params *resend.SendEmailRequest) (*resend.SendEmailResponse, error) {
		return nil, errCausePermanentFailure
	},
}
