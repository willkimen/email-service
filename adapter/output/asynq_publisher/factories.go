package asynqpublisher

import (
	"github.com/hibiken/asynq"
)

// NewAsynqEmailPublisherAdapter initializes and returns
// a new instance of AsynqEmailPublisherAdapter.
// It requires an active asynq.Client to manage
// the task queue.
func NewAsynqEmailPublisherAdapter(
	client *asynq.Client,
) *AsynqEmailPublisherAdapter {
	return &AsynqEmailPublisherAdapter{
		Client: client,
	}
}
