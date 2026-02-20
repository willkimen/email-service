package emailpublisher_test

import (
	"emailservice/adapter/output/asynq_publisher"
	"errors"
	"testing"

	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/require"
)

func TestPublish_Success(t *testing.T) {
	fake := &fakeEnqueuer{
		enqueueFunc: func(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
			return &asynq.TaskInfo{ID: "123"}, nil
		},
	}

	adapter := &emailpublisher.AsynqEmailPublisherAdapter{
		Client: fake,
	}

	msg := fakeMessage{}

	err := adapter.Publish(msg)

	require.NoError(
		t,
		err,
		"expected Publish to return nil error when enqueue succeeds",
	)
}

func TestPublish_EnqueueError(t *testing.T) {
	expectedErr := errors.New("enqueue failed")

	fake := &fakeEnqueuer{
		enqueueFunc: func(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
			return nil, expectedErr
		},
	}

	adapter := &emailpublisher.AsynqEmailPublisherAdapter{
		Client: fake,
	}

	msg := fakeMessage{}

	err := adapter.Publish(msg)

	require.Error(
		t,
		err,
		"expected Publish to return error when enqueue fails",
	)

	require.ErrorContains(
		t,
		err,
		"enqueue email task",
		"expected returned error to contain context message 'enqueue email task'",
	)

	require.ErrorIs(
		t,
		err,
		expectedErr,
		"expected returned error to wrap the original enqueue error",
	)
}
