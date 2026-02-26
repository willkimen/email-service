package emailpublisher

import (
	"emailservice/core/application/email_message"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/hibiken/asynq"
)

type Payload struct {
	To        string
	Subject   string
	EmailType string
	BodyData  any
}

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
// an EmailMessage into a background task. It is responsible only
// for serialization and task enqueueing, not for executing the email send itself.
type AsynqEmailPublisherAdapter struct {
	// Client is the Asynq client used to enqueue background tasks.
	Client taskEnqueuer
	Logger *slog.Logger
}

func NewAsynqEmailPublisherAdapter(
	client *asynq.Client,
	logger *slog.Logger,
) *AsynqEmailPublisherAdapter {
	return &AsynqEmailPublisherAdapter{
		Client: client,
		Logger: logger,
	}
}

// Publish serializes the given EmailMessage into a task payload
// and enqueues it for asynchronous processing.
//
// The method transforms the message into a transport-friendly
// structure, marshals it to JSON, and creates a task with a predefined
// type identifier.
//
// If serialization fails, a marshalling error is returned.
// If enqueueing fails, an infrastructure-level error is returned.
// The actual email delivery is expected to be handled by a separate worker.
func (a *AsynqEmailPublisherAdapter) Publish(message emailmessage.EmailMessage) error {
	a.Logger.Info(
		"publishing email task",
		"to", message.GetTo(),
		"subject", message.GetSubject(),
		"email_type", message.GetEmailType(),
	)

	payload, err := json.Marshal(Payload{
		To:        message.GetTo(),
		Subject:   message.GetSubject(),
		EmailType: message.GetEmailType(),
		BodyData:  message.GetBodyData(),
	})

	if err != nil {
		a.Logger.Error(
			"failed to marshal email payload",
			"error", err,
			"to", message.GetTo(),
			"subject", message.GetSubject(),
		)
		return fmt.Errorf("marshal email payload: %w", err)
	}

	task := asynq.NewTask("email:send", payload)

	info, err := a.Client.Enqueue(task)
	if err != nil {
		a.Logger.Error(
			"failed to enqueue email task",
			"error", err,
			"to", message.GetTo(),
			"subject", message.GetSubject(),
		)
		return fmt.Errorf("enqueue email task: %w", err)
	}

	a.Logger.Info(
		"email task enqueued successfully",
		"task_id", info.ID,
		"queue", info.Queue,
		"to", message.GetTo(),
		"subject", message.GetSubject(),
	)

	return nil
}
