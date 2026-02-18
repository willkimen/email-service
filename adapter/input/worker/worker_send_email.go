package worker

import (
	"context"
	"emailservice/core/application/email_errors"
	"emailservice/core/application/ports/input"
	"errors"

	"github.com/hibiken/asynq"
)

// SendEmailTaskHandler handles background tasks responsible for
// executing the email sending workflow.
//
// This struct belongs to the infrastructure layer and acts as an
// input adapter between Asynq and the application use case.
// It converts an incoming task into a domain EmailMessage and
// delegates execution to the core application.
//
// It is also responsible for translating domain-level errors
// into retry semantics understood by Asynq.
type SendEmailTaskHandler struct {
	// UseCase executes the synchronous email sending use case.
	UseCase inputport.ExecuteSendEmailInputPort
}

// ProcessSendEmail processes the "email:send" task received from Asynq.
//
// The method reconstructs the EmailMessage from the task payload
// and invokes the application use case.
//
// If payload deserialization fails, the task is considered invalid
// and retry is skipped.
//
// If the use case returns a temporary failure error, the error is
// propagated to allow Asynq to retry the task.
//
// Any non-temporary error is treated as a permanent failure and
// retry is skipped.
func (e *SendEmailTaskHandler) ProcessSendEmail(ctx context.Context, t *asynq.Task) error {
	message, err := ToEmailMessage(t.Payload())
	if err != nil {
		return asynq.SkipRetry
	}

	err = e.UseCase.ExecuteSend(message)
	if err != nil {
		if errors.Is(err, emailerrors.ErrTemporaryFailure) {
			return err
		}

		return asynq.SkipRetry
	}

	return nil
}
