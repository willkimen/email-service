package asynqpublisher

import (
	"emailservice/core/application/apperrors"
	"emailservice/core/application/message"
	"encoding/json"

	"github.com/hibiken/asynq"
)

// taskEnqueuer abstracts the enqueueing behavior required by
// AsynqEmailPublisherAdapter.
//
// This interface exists to decouple the adapter from the concrete
// *asynq.Client implementation, enabling easier testing and
// substitution of the underlying task queue mechanism.
//
// Any implementation must enqueue a task and return the associated
// TaskInfo or an error if the operation fails.
type taskEnqueuer interface {
	Enqueue(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error)
}

// AsynqEmailPublisherAdapter implements the output port responsible for
// publishing email sending requests using Asynq as the asynchronous mechanism.
//
// This adapter belongs to the infrastructure layer and translates
// an Message into a background task. It is responsible only
// for serialization and task enqueueing, not for executing the email send itself.
type AsynqEmailPublisherAdapter struct {
	// Client is the Asynq client used to enqueue background tasks.
	Client taskEnqueuer
}

// Publish serializes the given Message into a task payload
// and enqueues it for asynchronous processing.
//
// The method transforms the message into a transport-friendly
// structure, marshals it to JSON, and creates a task with a predefined
// type identifier.
//
// The actual email delivery is expected to be handled by a separate worker.
//
// Errors:
//   - Returns apperrors.InfrastructureError if the message fails to marshal into JSON.
//   - Returns apperrors.InfrastructureError if the background worker client fails
//     to enqueue the task.
func (a *AsynqEmailPublisherAdapter) Publish(message message.Message) error {
	task_payload, err := json.Marshal(message)

	if err != nil {
		return apperrors.NewInfrastructureError(
			"failed to serialize message in JSON",
			err,
		)

	}

	task := asynq.NewTask("email:send", task_payload)

	_, err = a.Client.Enqueue(task)
	if err != nil {
		return apperrors.NewInfrastructureError(
			"failed to enqueue message",
			err,
		)
	}

	return nil
}
