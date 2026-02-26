package worker

import (
	"context"
	"emailservice/core/application/email_errors"
	"emailservice/core/application/ports/input"
	"errors"
	"log/slog"

	"github.com/hibiken/asynq"
)

// SendEmailTaskHandler handles background tasks responsible for
// executing the email sending workflow.
//
// This struct belongs to the infrastructure layer and acts as an
// input adapter between Asynq and the application use case.
// It converts an incoming task into an EmailMessage and
// delegates execution to the core application.
type SendEmailTaskHandler struct {
	// UseCase executes the synchronous email sending use case.
	UseCase inputport.ExecuteSendEmailInputPort
	Logger  *slog.Logger
}

func NewSendEmailTaskHandler(
	usecase inputport.ExecuteSendEmailInputPort,
	logger *slog.Logger,
) *SendEmailTaskHandler {
	return &SendEmailTaskHandler{
		UseCase: usecase,
		Logger:  logger,
	}
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
	e.Logger.InfoContext(
		ctx,
		"processing email send task",
		"task_type", t.Type(),
	)

	message, err := ToEmailMessage(t.Payload())
	if err != nil {
		e.Logger.ErrorContext(
			ctx,
			"failed to deserialize email task payload",
			"error", err,
			"task_type", t.Type(),
		)
		return asynq.SkipRetry
	}

	err = e.UseCase.ExecuteSend(message)
	if err != nil {
		if errors.Is(err, emailerrors.ErrTemporaryFailure) {
			e.Logger.ErrorContext(
				ctx,
				"temporary failure while sending email, will retry",
				"error", err,
				"task_type", t.Type(),
				"to", message.GetTo(),
				"subject", message.GetSubject(),
			)
			return err
		}

		e.Logger.ErrorContext(
			ctx,
			"permanent failure while sending email, skipping retry",
			"error", err,
			"task_type", t.Type(),
			"to", message.GetTo(),
			"subject", message.GetSubject(),
		)
		return asynq.SkipRetry
	}

	e.Logger.InfoContext(
		ctx,
		"email task processed successfully",
		"task_type", t.Type(),
		"to", message.GetTo(),
		"subject", message.GetSubject(),
	)

	return nil
}
