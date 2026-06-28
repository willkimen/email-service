package worker

import (
	"context"
	"emailservice/core/application/apperrors"
	"emailservice/core/application/message"
	"emailservice/core/application/ports/input"
	"encoding/json"
	"errors"
	"log/slog"

	"github.com/hibiken/asynq"
)

// SendEmailTaskHandler handles background tasks responsible for
// executing the email sending workflow.
//
// It acts as an input adapter between Asynq and the application use case.
// It converts an incoming task into a Message and
// delegates execution to the core application.
type SendEmailTaskHandler struct {
	// UseCase executes the synchronous email sending use case.
	UseCase inputport.ExecuteSendEmailPort

	// Logger provides standard capabilities to log information, warnings,
	// and execution errors during lifecycle.
	Logger *slog.Logger
}

// ProcessSendEmail processes the "email:send" task received from Asynq.
//
// The method reconstructs the Message from the task payload
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
	taskID, _ := asynq.GetTaskID(ctx)
	queue, _ := asynq.GetQueueName(ctx)
	retryCount, _ := asynq.GetRetryCount(ctx)

	e.Logger.InfoContext(
		ctx,
		"processing email send task",
		"task_id", taskID,
		"task_type", t.Type(),
		"queue", queue,
		"retry_count", retryCount,
	)

	var message message.Message
	err := json.Unmarshal(t.Payload(), &message)

	if err != nil {
		e.Logger.ErrorContext(
			ctx,
			"failed to deserialize email task payload",
			"task_type", t.Type(),
			"error", err,
		)
		return asynq.SkipRetry
	}

	err = e.UseCase.Execute(message)
	if err != nil {
		if errors.Is(err, apperrors.ErrTemporaryFailure) {
			e.Logger.ErrorContext(
				ctx,
				"temporary failure while sending email, will retry",
				"task_type", t.Type(),
				"error", err,
				"to", message.To,
				"message_id", message.Id,
			)
			return err
		}

		e.Logger.ErrorContext(
			ctx,
			"permanent failure while sending email, skipping retry",
			"task_type", t.Type(),
			"error", err,
			"to", message.To,
			"message_id", message.Id,
		)
		return asynq.SkipRetry
	}

	return nil
}
